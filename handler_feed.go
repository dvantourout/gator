package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
	"github.com/dvantourout/gator/internal/rss"
)

func aggHandler(s *state, cmd command) error {
	rss, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("cannot fetch rss: %w", err)
	}

	fmt.Println(rss)

	return nil
}

func addfeed(s *state, cmd command) error {
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
