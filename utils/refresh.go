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

func newError(s string, i ...interface{}) error {
	return errors.New(fmt.Sprintf(s, i...))
}

func (s Store) Add(link string) error {
	r, err := http.Get(link)
	if err != nil {
		return newError("Could not reach %s.", link)
	}
	defer r.Body.Close()
	xml, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return newError("Could not get XML feed for this channel from %s.", link)
	}
	channel, err := ParseFeed(string(xml))
	if err != nil {
		return err
	}
	path := filepath.Join(s.RootFolder, channel.Title)
	_, err = os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		return newError("Channel by the same title (%s) already exists.", channel.Title)
	}
	err = os.Mkdir(path, 0750)
	if err != nil {
		return newError("Could not create folder %s.", path)
	}
	feed := filepath.Join(path, s.FeedName)
	err = os.WriteFile(feed, xml, 0666)
	if err != nil {
		return newError("Could not write to feed file %s.", feed)
	}
	return s.Refresh(channel.Title, 5, true)
}

func (s Store) Refresh(title string, length int, overwrite bool) error {
	channel, err := ParseFile(filepath.Join(s.RootFolder, title, s.FeedName))
	if err != nil {
		return err
	}

	newFeed, err := http.Get(channel.FeedURL)
	if err != nil {
		return newError("Could not retrieve new feed at %s", channel.FeedURL)
	}
	feedContents, err := ioutil.ReadAll(newFeed.Body)
	if err != nil {
		return newError("Could not read response from %s", channel.FeedURL)
	}
	newFeed.Body.Close()
	fp := filepath.Join(s.RootFolder, title, s.FeedName)
	err = os.WriteFile(fp, feedContents, 0666)
	if err != nil {
		return newError("Could not write to file %s", fp)
	}

	channel, err = ParseFeed(string(feedContents))
	if err != nil {
		return err
	}

	for _, item := range channel.Items[0:length] {
		_, err := os.Stat(fmt.Sprintf("%s-%s", item.Title, item.PubDate))
		if overwrite || errors.Is(err, os.ErrNotExist) {
			r, err := http.Get(item.AV.Url)
			if err != nil {
				return newError("Could not retrieve episode at %s", item.AV.Url)
			}
			cast, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return newError("Could not read response from %s", item.AV.Url)
			}
			r.Body.Close()
			fp := filepath.Join(s.RootFolder, title, item.FileName())
			err = os.WriteFile(fp, cast, 0666)
			if err != nil {
				return newError("Could not write to file %s", fp)
			}
		}
	}

	return nil
}

func (s Store) RefreshAll(length int) error {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return newError("Could not access %s", s.RootFolder)
	}
	for _, l := range list {
		err := s.Refresh(l.Name(), length, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Store) Remove(title string) error {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return newError("Could not access %s", s.RootFolder)
	}
	exists := false
	for _, l := range list {
		if l.Name() == title {
			exists = true
		}
	}
	if exists {
		path := filepath.Join(s.RootFolder, title)
		err := os.RemoveAll(path)
		if err != nil {
			return newError("Could not remove %s", path)
		}
		return nil
	}
	return newError("Specified channel (%s) does not exist.", title)
}
