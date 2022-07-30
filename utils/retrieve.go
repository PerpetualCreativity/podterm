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

func (s Store) GetEpisodePath(channel string, index int) (string, error) {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return "", newError("Could not access %s", s.RootFolder)
	}
	for _, l := range list {
		if l.IsDir() && l.Name() == channel {
			path := filepath.Join(s.RootFolder, l.Name(), s.FeedName)
			xml, err := os.ReadFile(path)
			if err != nil {
				return "", newError("Could not access %s", path)
			}
			ch, err := ParseFeed(string(xml))
			if err != nil {
				return "", err
			}
			if index > len(ch.Items) {
				return "", newError("The %s episode does not exist.", humanize.Ordinal(index))
			}
			episode := ch.Items[index]
			path = filepath.Join(s.RootFolder, l.Name(), episode.FileName())
			_, err = os.Stat(path)
			if errors.Is(err, os.ErrNotExist) {
				return "", newError("The %s episode has not been downloaded", humanize.Ordinal(index))
			}
			return path, nil
		}
	}
	return "", newError("The specified channel (%s) could not be found.", channel)
}
