package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func (s Store) ChannelList() ([]string, error) {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return nil, fmt.Errorf("could not access %s", s.RootFolder)
	}
	var filtered []string
	for _, l := range list {
		if l.IsDir() {
			filtered = append(filtered, l.Name())
		}
	}
	return filtered, nil
}

func (s Store) FindChannel(search string) (result string, options []string, err error) {
	in, err := s.ChannelList()
	if err != nil { return "", nil, err }
	f := fuzzy.RankFindNormalizedFold(search, in)
	if len(f) == 0 {
		return "", nil, fmt.Errorf("no matches for %s found", search)
	}
	if len(f) > 1 {
		sort.Sort(f)
	}
	for _, t := range f {
		options = append(options, t.Target)
	}
	result = f[0].Target
	return
}

func (s Store) DownloadedEpisodeList(channel string) ([]Item, error) {
	list, err := os.ReadDir(filepath.Join(s.RootFolder, channel))
	if err != nil {
		return nil, fmt.Errorf("could not access %s", s.RootFolder)
	}
	ch, err := ParseFile(filepath.Join(s.RootFolder, channel, s.FeedName))
	if err != nil { return nil, err }
	downloaded := make([]Item, 0, len(list))
	for _, episode := range ch.Items {
		for _, l := range list {
			if strings.Split(l.Name(), ".")[0] != episode.Guid {
				downloaded = append(downloaded, episode)
			}
		}
	}
	return downloaded, nil
}

func (s Store) GetEpisode(channel string, index int, overwrite bool) (string, error) {
	ch, err := ParseFile(filepath.Join(s.RootFolder, channel, s.FeedName))
	if err != nil { return "", err }
	if index > len(ch.Items) {
		return "", fmt.Errorf("there are only %d episodes in this channel", index)
	}
	episode := ch.Items[index]
	path := filepath.Join(s.RootFolder, channel, episode.FileName())
	_, err = os.Stat(path)
	if overwrite || errors.Is(err, os.ErrNotExist) {
		r, err := http.Get(episode.AV.Url)
		if err != nil {
			return "", fmt.Errorf("could not retrieve episode at %s", episode.AV.Url)
		}
		cast, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return "", fmt.Errorf("could not read response from %s", episode.AV.Url)
		}
		err = r.Body.Close()
		if err != nil {
			return "", fmt.Errorf("could not read response from %s", episode.AV.Url)
		}
		fp := filepath.Join(s.RootFolder, channel, episode.FileName())
		err = os.WriteFile(fp, cast, 0666)
		if err != nil {
			return "", fmt.Errorf("could not write to file %s", fp)
		}
	}

	return path, nil
}
