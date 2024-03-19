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

func main() {
	flag.Parse()

	if *url == "" {
		fmt.Fprintln(os.Stderr, "url must be provided")
		os.Exit(2)
	}

	res, err := http.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong fetching url: %s\n", *url)
		os.Exit(1)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	atom := new(Atom)
	err = xml.Unmarshal(data, &atom)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(atom.Entries[0].Content.Value)
}
