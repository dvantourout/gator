package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("add feed should be called with 2 arguments")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	now := time.Now().UTC()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		UpdatedAt: now,
		CreatedAt: now,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return nil
	}

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Printing feeds:")
	for _, feed := range feeds {
		fmt.Printf("* Name: %s User: %s Url: %s\n", feed.Name, feed.UserName, feed.Url)
	}

	return nil
}
