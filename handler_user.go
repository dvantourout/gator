package main

import "fmt"

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("loginHandler need one argument")
	}

	username := cmd.args[0]
	if err := s.config.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user has been set")
	return nil
}
