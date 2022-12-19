package types

import "encoding/xml"

type GeophysicsRss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Text          string  `xml:",chardata"`
	Title         string  `xml:"title"`
	Link          string  `xml:"link"`
	Description   string  `xml:"description"`
	PubDate       string  `xml:"pubDate"`
	LastBuildDate string  `xml:"lastBuildDate"`
	WebMaster     string  `xml:"webMaster"`
	Items         []Items `xml:"item"`
}

type Items struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
}
