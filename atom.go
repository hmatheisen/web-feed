package main

import "encoding/xml"

type Atom struct {
	XMLName     xml.Name  `xml:"feed"`
	Author      Person    `xml:"author"`
	Contributor Person    `xml:"contributor"`
	Generator   Generator `xml:"generator"`
	Icon        string    `xml:"icon"`
	ID          string    `xml:"id"`
	Link        Link      `xml:"link"`
	Logo        string    `xml:"logo"`
	Rights      string    `xml:"rights"`
	Subtitle    string    `xml:"subtitle"`
	Title       string    `xml:"title"`
	Updated     string    `xml:"updated"`
	Entries     []Entry   `xml:"entry"`
}

type Person struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URI   string `xml:"uri"`
}

type Generator struct {
	URI     string `xml:"uri,attr"`
	Version string `xml:"version,attr"`
}

type Link struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	HrefLang string `xml:"hreflang,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}

type Entry struct {
	Author      Person   `xml:"author"`
	Category    Category `xml:"category"`
	Content     Content  `xml:"content"`
	Contributor Person   `xml:"contributor"`
	ID          string   `xml:"id"`
	Link        Link     `xml:"link"`
	Published   string   `xml:"published"`
	Rights      string   `xml:"rights"`
	Source      string   `xml:"source"`
	Summary     string   `xml:"summary"`
	Title       string   `xml:"title"`
	Updated     string   `xml:"updated"`
}

type Category struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr"`
	Label  string `xml:"label,attr"`
}

type Content struct {
	Type  string `xml:"type,attr"`
	Src   string `xml:"src,attr"`
	Value string `xml:",innerxml"`
}

func (a Atom) List(count int) []Article {
	// make count shorter id there are less articles
	if len(a.Entries) < count {
		count = len(a.Entries)
	}

	articles := make([]Article, count)
	entries := a.Entries

	for i := 0; i < count; i++ {
		articles[i].Title = entries[i].Title
		articles[i].Description = entries[i].Content.Value
		articles[i].Link = entries[i].Link.Href
	}

	return articles
}
