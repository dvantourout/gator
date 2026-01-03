package main

import (
	"context"
	"fmt"

	"github.com/dvantourout/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	rss, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("cannot fetch rss: %w", err)
	}

	fmt.Println(rss)

	return nil
}
