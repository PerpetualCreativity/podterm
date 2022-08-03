package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
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

func (s Store) EpisodeList(channel string) ([]string, error) {
	path := filepath.Join(s.RootFolder, channel)
	list, err := os.ReadDir(path)
	if err != nil { return nil, newError("Could not access %s", path) }
	var filtered []string
	for _, l := range list {
		if l.Name() != s.FeedName {
			filtered = append(filtered, l.Name())
		}
	}
	return filtered, nil
}

func (s Store) GetEpisodePath(channel string, index int) (string, error) {
	ch, err := ParseFile(filepath.Join(s.RootFolder, channel, s.FeedName))
	if err != nil { return "", err }
	if index > len(ch.Items) {
		return "", newError("The %s episode does not exist.", humanize.Ordinal(index))
	}
	episode := ch.Items[index]
	path := filepath.Join(s.RootFolder, channel, episode.FileName())
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", newError("The %s episode has not been downloaded", humanize.Ordinal(index))
	}
	return path, nil
}
