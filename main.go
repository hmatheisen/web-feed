package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
)

var (
	url       = flag.String("url", "", "rss url")
	count     = flag.Int("count", 5, "articles count")
	save      = flag.Bool("save", false, "save a feed url")
	listSaved = flag.Bool("list-saved", false, "list articles from saved urls")
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

func usage() {
	fmt.Fprintln(
		os.Stderr,
		"usage: web-feed [--url=<url> [--save] [--count=<count>]] | [--list-saved]",
	)
	os.Exit(1)
}

func fetchSavedUrls() ([]string, error) {
	var urls []string

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path.Join(home, DBDir, DBFile))
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls, nil
}

func listArticlesFromSavedUrls(count int) error {
	var wg sync.WaitGroup

	urls, err := fetchSavedUrls()
	if err != nil {
		return err
	}

	feeds := make(map[string]Feed)
	for _, url := range urls {
		wg.Add(1)

		go func(url *string, count int) {
			defer wg.Done()

			feed, _ := NewFeed(url)
			feeds[*url] = feed

		}(&url, count)
	}

	wg.Wait()

	for url, feed := range feeds {
		fmt.Println(url)
		for _, article := range feed.List(count) {
			fmt.Println(article)
		}
	}

	return nil
}

func main() {
	flag.Parse()
	flag.Usage = usage

	if *url == "" && !*listSaved {
		flag.Usage()
	}

	if *listSaved && *save {
		flag.Usage()
	}

	if *listSaved {
		err := listArticlesFromSavedUrls(*count)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return
	}

	if *save {
		go saveUrl(*url)
	}

	feed, err := NewFeed(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(*url)
	for _, article := range feed.List(*count) {
		fmt.Println(article)
	}
}
