package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
	"github.com/dvantourout/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("agg take a least one argument")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feedData, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	for _, item := range feedData.Chanel.Items {
		now := time.Now().UTC()

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			UpdatedAt:   now,
			CreatedAt:   now,
			Title:       item.Title,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
			Url:         item.Link,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}
