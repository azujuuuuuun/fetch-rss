package main

import (
	"encoding/xml"
	"fmt"
	"io"
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

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("URL arg is only required.")
		return
	}
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Fetching RSS failed. %#v", err)
		return
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		fmt.Printf("io.ReadAll error occurred. %#v", err)
		return
	}

	var rss RSS
	if err := xml.Unmarshal(b, &rss); err != nil {
		fmt.Printf("Parsing XML Failed. %#v", err)
		return
	}

	// TODO: Format output shape
	fmt.Printf("%#v", rss)
}
