package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Article struct {
	Title string
	Link  string
}

func (a Article) String() string {
	return fmt.Sprintf("Title: %s\nLink: %s\n", a.Title, a.Link)
}

type Feed interface {
	List(count int) []Article
}

func detectFeedType(data []byte) (*string, error) {
	var feed struct {
		XMLName xml.Name
	}

	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return &feed.XMLName.Local, nil
}

func NewFeed(url *string) (Feed, error) {
	var feed Feed

	res, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	feedType, err := detectFeedType(data)
	if err != nil {
		return nil, err
	}

	switch *feedType {
	default:
		return nil, fmt.Errorf("Unknown feed type: %s\n", *feedType)
	case "feed":
		feed = new(Atom)
	case "rss":
		feed = new(RSS)
	}

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
