package utils

import (
	"encoding/xml"
	"fmt"
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
	Duration    string `xml:"itunes:duration"`
	Explicit    string `xml:"itunes:explicit"`
	Subtitle    string `xml:"itunes:subtitle"`
	Description string `xml:"itunes:description"`
	Link        string `xml:"link"`
	AV          AV     `xml:"enclosure"`
}
func (item Item) FileName() string {
	return fmt.Sprintf("%s-%s.%s", item.Title, item.PubDate, item.AV.SmartType())
}
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Link        string `xml:"link"`
	Copyright   string `xml:"copyright"`
	Explicit    string `xml:"itunes:explicit"`
	Image       Image  `xml:"image"`
	Items       []Item `xml:"item"`
}

func ParseFeed(xmlSource string) (Channel, error) {
	type Result struct {
		Channel Channel `xml:"channel"`
	}
	v := Result{}

	err := xml.Unmarshal([]byte(xmlSource), &v)
	return v.Channel, err
}

