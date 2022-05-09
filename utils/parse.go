package parse

import (
	"encoding/xml"
)
type Image struct {
	Link  string `xml:"link"`
	Url   string `xml:"url"`
	Title string `xml:"title"`
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
}
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Link        string `xml:"link"`
	Copyright   string `xml:"copyright"`
	Explicit    string `xml:"itunes:explicit"`
	Image       Image  `xml:"image"`
	FeedURL     string `xml:"itunes:new-feed-url"`
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

