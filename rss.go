package main

type RSS struct {
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Author      string `xml:"author"`
	Category    string `xml:"category"`
	Comments    string `xml:"comments"`
	Enclosure   string `xml:"enclosure"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Source      string `xml:"source"`
}

func (rss RSS) List(count int) any {
	// make count shorter id there are less articles
	if len(rss.Channel.Items) < count {
		count = len(rss.Channel.Items)
	}

	articles := make([]Article, count)
	items := rss.Channel.Items

	for i := 0; i < count; i++ {
		articles[i].Title = items[i].Title
		articles[i].Description = items[i].Description
		articles[i].Link = items[i].Link
	}

	return articles
}
