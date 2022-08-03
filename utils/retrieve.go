package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (s Store) ChannelList() ([]string, error) {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return nil, newError("Could not access %s", s.RootFolder)
	}
	var filtered []string
	for _, l := range list {
		if l.IsDir() {
			filtered = append(filtered, l.Name())
		}
	}
	return filtered, nil
}

func (s Store) GetEpisode(channel string, index int, overwrite bool) (string, error) {
	ch, err := ParseFile(filepath.Join(s.RootFolder, channel, s.FeedName))
	if err != nil { return "", err }
	if index > len(ch.Items) {
		return "", newError("There are only %d episodes in this channel.", index)
	}
	episode := ch.Items[index]
	path := filepath.Join(s.RootFolder, channel, episode.FileName())
	_, err = os.Stat(path)
	if overwrite || errors.Is(err, os.ErrNotExist) {
		r, err := http.Get(episode.AV.Url)
		if err != nil {
			return "", newError("Could not retrieve episode at %s", episode.AV.Url)
		}
		cast, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return "", newError("Could not read response from %s", episode.AV.Url)
		}
		err = r.Body.Close()
		if err != nil {
			return "", newError("Could not read response from %s", episode.AV.Url)
		}
		fp := filepath.Join(s.RootFolder, channel, episode.FileName())
		err = os.WriteFile(fp, cast, 0666)
		if err != nil {
			return "", newError("Could not write to file %s", fp)
		}
	}

	return path, nil
}
