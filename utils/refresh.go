package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Store struct {
	RootFolder string
	FeedName   string
}

func (s Store) Add(link string) error {
	r, err := http.Get(link)
	if err != nil {
		return fmt.Errorf("could not reach %s", link)
	}
	defer r.Body.Close()
	xml, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not get XML feed for this channel from %s", link)
	}
	channel, err := ParseFeed(string(xml))
	if err != nil {
		return err
	}
	path := filepath.Join(s.RootFolder, channel.Title)
	_, err = os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("channel by the same title (%s) already exists", channel.Title)
	}
	err = os.Mkdir(path, 0750)
	if err != nil {
		return fmt.Errorf("could not create folder %s", path)
	}
	feed := filepath.Join(path, s.FeedName)
	err = os.WriteFile(feed, xml, 0666)
	if err != nil {
		return fmt.Errorf("could not write to feed file %s", feed)
	}
	return nil
}

func (s Store) Refresh(title string) error {
	channel, err := ParseFile(filepath.Join(s.RootFolder, title, s.FeedName))
	if err != nil {
		return err
	}

	newFeed, err := http.Get(channel.FeedURL)
	if err != nil {
		return fmt.Errorf("could not retrieve new feed at %s", channel.FeedURL)
	}
	feedContents, err := ioutil.ReadAll(newFeed.Body)
	if err != nil {
		return fmt.Errorf("could not read response from %s", channel.FeedURL)
	}
	err = newFeed.Body.Close()
	if err != nil {
		return fmt.Errorf("could not read response from %s", channel.FeedURL)
	}
	fp := filepath.Join(s.RootFolder, title, s.FeedName)
	err = os.WriteFile(fp, feedContents, 0666)
	if err != nil {
		return fmt.Errorf("could not write to file %s", fp)
	}

	return nil
}

func (s Store) RefreshAll() error {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return fmt.Errorf("could not access %s", s.RootFolder)
	}
	for _, l := range list {
		err = s.Refresh(l.Name())
		if err != nil {
			return err
		}
	}
	return nil
}
