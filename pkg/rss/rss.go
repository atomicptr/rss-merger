package rss

type Rss struct {
	XMLName string  `xml:"rss"`
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	XMLName       string `xml:"channel"`
	Title         string `xml:"title"`
	Description   string `xml:"description"`
	Link          string `xml:"link,omitempty"`
	LastBuildDate string `xml:"lastBuildDate"`
	Item          []Item `xml:"item"`
}

type Item struct {
	Title   string `xml:"title,omitempty"`
	Link    string `xml:"link"`
	GUID    string `xml:"guid"`
	PubDate string `xml:"pubDate,omitempty"`
}

func New(title, description, link string) Rss {
	return Rss{
		Version: "2.0",
		Channel: Channel{
			Title:       title,
			Description: description,
			Link:        link,
		},
	}
}
