package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
	"github.com/google/uuid"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("login command need to be called with 1 argument")
	}

	username := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("user has been set")
	return nil
}

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("register command should be called with 1 arguments")
	}

	name := cmd.args[0]
	now := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), Name: name, UpdatedAt: now, CreatedAt: now})
	if err != nil {
		return err
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	return nil
}

func resetHandler(s *state, cmd command) error {
	return s.db.Reset(context.Background())
}
