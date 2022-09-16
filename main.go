package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RSS struct {
	Channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Image       struct {
			Url   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Generator     string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        string `xml:"guid"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchRSS(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request HTTP Get: %w", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read all response body: %w", err)
	}

	var rss RSS
	if err := xml.Unmarshal(b, &rss); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &rss, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("URL arg is only required.")
		return
	}
	url := os.Args[1]

	rss, err := fetchRSS(url)
	if err != nil {
		log.Fatalf("failed to fetching RSS: %v", err)
		return
	}

	// TODO: Format output shape
	fmt.Printf("%#v", rss)
}
