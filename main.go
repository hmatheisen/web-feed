package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var url = flag.String("url", "", "rss url")

type Article struct {
	Title       string
	Link        string
	Description string
}

func (a Article) String() string {
	return fmt.Sprintf(
		"Title: %s\nLink: %s\nDescription: %s\n",
		a.Title,
		a.Link,
		a.Description,
	)
}

type Feed interface {
	List(count int) []Article
}

func NewFeed(url *string) (Feed, error) {
	// TODO: find a way to differenciate RSS from Atom
	feed := new(Atom)

	res, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func main() {
	flag.Parse()
	if *url == "" {
		fmt.Fprintln(os.Stderr, "url must be provided")
		os.Exit(2)
	}

	feed, err := NewFeed(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, article := range feed.List(3) {
		fmt.Println(article)
	}
}
