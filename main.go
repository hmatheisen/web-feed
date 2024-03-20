package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	url   = flag.String("url", "", "rss url")
	count = flag.Int("count", 5, "articles count")
	save  = flag.Bool("save", false, "save a feed url")
)

const (
	DBDir  = ".web-feed"
	DBFile = "feeds"
)

func saveUrl(url string) {
	var sb strings.Builder
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = os.Mkdir(path.Join(home, DBDir), 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file, err := os.OpenFile(
		path.Join(home, DBDir, DBFile),
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == url {
			return
		}
	}

	sb.WriteString(url)
	sb.WriteRune('\n')

	_, err = file.WriteString(sb.String())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func main() {
	flag.Parse()
	if *url == "" {
		fmt.Fprintln(os.Stderr, "url must be provided")
		os.Exit(2)
	}

	if *save {
		go saveUrl(*url)
	}

	feed, err := NewFeed(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, article := range feed.List(*count) {
		fmt.Println(article)
	}
}
