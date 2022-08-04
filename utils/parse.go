package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Image struct {
	Link  string `xml:"link"`
	Url   string `xml:"url"`
	Title string `xml:"title"`
}
type AV struct {
	Url    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length string `xml:"length,attr"`
}

func (av AV) SmartType() string {
	tb, ta, f := strings.Cut(av.Type, "/")
	if f {
		return ta
	} else {
		return tb
	}
}

type Item struct {
	Title       string `xml:"title"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Duration    string `xml:"duration"`
	Explicit    string `xml:"explicit"`
	Subtitle    string `xml:"subtitle"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	AV          AV     `xml:"enclosure"`
}

func (item Item) FileName() string {
	return fmt.Sprintf("%s.%s", item.Guid, item.AV.SmartType())
}

type Channel struct {
	FeedURL struct {
		Href string `xml:"href,attr"`
	} `xml:"atom link"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Link        string `xml:"_ link"`
	Copyright   string `xml:"copyright"`
	Image       Image  `xml:"image"`
	Items       []Item `xml:"item"`
}

func ParseFeed(xmlSource string) (Channel, error) {
	type Result struct {
		Channel Channel `xml:"channel"`
	}
	v := Result{}

	xmlSource = strings.ReplaceAll(xmlSource, " xmlns:atom=\"http://www.w3.org/2005/Atom\"", "")

	decoder := xml.NewDecoder(strings.NewReader(xmlSource))
	decoder.DefaultSpace = "_"

	err := decoder.Decode(&v)

	return v.Channel, err
}

func ParseFile(path string) (Channel, error) {
	xml, err := os.ReadFile(path)
	if err != nil {
		return Channel{}, fmt.Errorf("could not access %s", path)
	}
	feed, err := ParseFeed(string(xml))
	if err != nil {
		return Channel{}, err
	}
	return feed, nil
}