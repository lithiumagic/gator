package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil) // creates a *http.Request object.
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	feedRSS := RSSFeed{}
	err = xml.Unmarshal(bytes, &feedRSS)
	if err != nil {
		return nil, err
	}
	feedRSS.Channel.Title = html.UnescapeString(feedRSS.Channel.Title)
	feedRSS.Channel.Description = html.UnescapeString(feedRSS.Channel.Description)
	for i := range feedRSS.Channel.Item {
		feedRSS.Channel.Item[i].Title = html.UnescapeString(feedRSS.Channel.Item[i].Title)
		feedRSS.Channel.Item[i].Description = html.UnescapeString(feedRSS.Channel.Item[i].Description)
	}

	return &feedRSS, nil
}
