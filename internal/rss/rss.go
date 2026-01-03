package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Chanel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"Link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	rssFeed.Chanel.Title = html.UnescapeString(rssFeed.Chanel.Title)
	rssFeed.Chanel.Description = html.UnescapeString(rssFeed.Chanel.Description)
	for i, item := range rssFeed.Chanel.Items {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Chanel.Items[i] = item
	}

	return &rssFeed, nil
}
