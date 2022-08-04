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
	xml, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not get XML feed for this channel from %s", link)
	}
	err = r.Body.Close()
	if err != nil {
		return fmt.Errorf("could not read feed from %s")
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

func (s Store) Refresh(title string) ([]Item, error) {
	channel, err := ParseFile(filepath.Join(s.RootFolder, title, s.FeedName))
	if err != nil {
		return nil, err
	}
	oldTop := ""
	if channel.Items != nil && len(channel.Items) != 0 {
		oldTop = channel.Items[0].Guid
	}

	newFeed, err := http.Get(channel.FeedURL)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve new feed at %s", channel.FeedURL)
	}
	feedContents, err := ioutil.ReadAll(newFeed.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response from %s", channel.FeedURL)
	}
	err = newFeed.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not read response from %s", channel.FeedURL)
	}
	fp := filepath.Join(s.RootFolder, title, s.FeedName)
	err = os.WriteFile(fp, feedContents, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not write to file %s", fp)
	}

	channel, err = ParseFeed(string(feedContents))
	if err != nil { return nil, err }
	var newItems []Item
	for i:=0; channel.Items[i].Guid!=oldTop; i++ {
		newItems = append(newItems, channel.Items[i])
	}

	return newItems, nil
}

type Collection struct {
	Channel  string
	Episodes []Item
}

func (s Store) RefreshAll() ([]Collection, error) {
	list, err := os.ReadDir(s.RootFolder)
	if err != nil {
		return nil, fmt.Errorf("could not access %s", s.RootFolder)
	}
	newCollections := make([]Collection, len(list))
	for i, l := range list {
		ni, err := s.Refresh(l.Name())
		if err != nil {
			return nil, err
		}
		newCollections[i] = Collection{
			Channel: l.Name(),
			Episodes: ni,
		}
	}
	return newCollections, nil
}
