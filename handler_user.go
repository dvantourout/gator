package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvantourout/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
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

func handlerRegister(s *state, cmd command) error {
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

func handlerReset(s *state, cmd command) error {
	return s.db.Reset(context.Background())
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("too many arguments")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Print("* " + user.Name)
		if s.config.CurrentUserName == user.Name {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}
