package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type XMLFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Channel XMLChannel `xml:"channel"`
}

type XMLChannel struct {
	Title         string     `xml:"title"`
	Link          string     `xml:"link"`
	Description   string     `xml:"description"`
	Language      string     `xml:"language"`
	Generator     string     `xml:"generator"`
	LastBuildDate CustomTime `xml:"lastBuildDate"`
	Item          []XMLItem  `xml:"item"`
}

type XMLItem struct {
	Title        string     `xml:"title"`
	Link         string     `xml:"link"`
	Published    CustomTime `xml:"pubDate"` // Custom time type for pubDate
	ReferenceUrl string     `xml:"guid"`
	Description  string     `xml:"description"`
}

type CustomTime struct {
	time.Time
}
type FeedRss struct {
	Title       string     `json:"title"`
	Link        string     `json:"link"`
	PublishedAt CustomTime `json:"publishedAt"`
	Description string     `json:"description"`
}

func (ct *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	dateFormats := []string{
		"Mon, 02 Jan 2006 15:04:05 -0700", // RFC 1123 (common in RSS feeds)
		"2006-01-02T15:04:05Z07:00",       // RFC 3339 (ISO 8601 format)
		"Mon, 02 Jan 2006 15:04:05 +0000", // RFC 1123 (with "+0000" for UTC)
		"2006-01-02 15:04:05",             // Common MySQL format
	}

	var parsedTime time.Time
	var err error
	for _, format := range dateFormats {
		parsedTime, err = time.Parse(format, raw)
		if err == nil {
			ct.Time = parsedTime
			return nil
		}
	}
	return nil
}

func GetXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("GET error:", err)
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read body error:", err)
		return "", fmt.Errorf("read body: %v", err)
	}
	return string(data), nil
}

func ParseRssXML(url string) ([]FeedRss, error) {
	data, err := GetXML(url)
	if err != nil {
		log.Printf("Error fetching XML: %v", err)
		return nil, err
	}

	var xmlFeed XMLFeed
	bytearray := []byte(data)

	err = xml.Unmarshal(bytearray, &xmlFeed)
	if err != nil {
		log.Printf("Error unmarshalling XML: %v", err)
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}
	var feedRss []FeedRss
	for _, feed := range xmlFeed.Channel.Item {
		rssFeed := FeedRss{
			Title:       feed.Title,
			Link:        feed.Link,
			PublishedAt: feed.Published,
			Description: feed.Description,
		}
		feedRss = append(feedRss, rssFeed)
	}
	return feedRss, nil
}
