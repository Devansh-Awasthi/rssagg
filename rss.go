package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// RSS is the top-level structure for the entire RSS feed.
// It corresponds to the <rss> element.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel contains the metadata for the feed and a list of news items.
// It corresponds to the <channel> element.
type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Language      string   `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Copyright     string   `xml:"copyright"`
	Image         Image    `xml:"image"`
	Items         []Item   `xml:"item"` // A slice to hold multiple <item> elements
}

// Image represents the feed's associated image.
// It corresponds to the <image> element.
type Image struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Link  string `xml:"link"`
}

// Item represents a single news article or entry in the feed.
// It corresponds to the <item> element.
type Item struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	GUID        string    `xml:"guid"`
	PubDate     string    `xml:"pubDate"`
	Enclosure   Enclosure `xml:"enclosure"`
}

// Enclosure represents media content attached to an item.
// Note how its fields are mapped from attributes, not elements.
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}
func urlTofeed(url string) (RSS,error){
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	res,err := httpClient.Get(url)
	if err!=nil {
		return RSS{},err
	}
	defer res.Body.Close()
	dat,err := io.ReadAll(res.Body)
	if err!=nil {
		return RSS{},err
	}
	rssFeed := RSS{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSS{}, err
	}
	return rssFeed, nil
}