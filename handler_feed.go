package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("add feed should be called with 2 arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]
	now := time.Now().UTC()

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		UpdatedAt: now,
		CreatedAt: now,
	})
	if err != nil {
		return err
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
		fmt.Printf("* Name: %s User: %s Url: %s\n", feed.Name, feed.UserName.String, feed.Url)
	}

	return nil
}
