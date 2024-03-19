package main

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      string   `xml:"author"`
	Category    string   `xml:"category"`
	Comments    string   `xml:"comments"`
	Enclosure   string   `xml:"enclosure"`
	GUID        string   `xml:"guid"`
	PubDate     string   `xml:"pubDate"`
	Source      string   `xml:"source"`
}
