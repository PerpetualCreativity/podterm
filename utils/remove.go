package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s Store) Clear(channel string) error {
	path := filepath.Join(s.RootFolder, channel)
	list, err := os.ReadDir(path)
	if err != nil { return fmt.Errorf("could not read %s", path) }
	for _, l := range list {
		if l.Name() != s.FeedName {
			episodePath := filepath.Join(path, l.Name())
			err = os.Remove(episodePath)
			if err != nil { return fmt.Errorf("could not remove %s", episodePath) }
		}
	}
	return nil
}

func (s Store) ClearAll() error {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil { return fmt.Errorf("could not read %s", s.RootFolder) }
	for _, l := range list {
		err = s.Clear(l.Name())
		if err != nil { return err }
	}
	return nil
}

func (s Store) Remove(channel string) error {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return fmt.Errorf("could not access %s", s.RootFolder)
	}
	exists := false
	for _, l := range list {
		if l.Name() == channel {
			exists = true
		}
	}
	if exists {
		path := filepath.Join(s.RootFolder, channel)
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("could not remove %s", path)
		}
		return nil
	}
	return fmt.Errorf("specified channel (%s) does not exist", channel)
}
