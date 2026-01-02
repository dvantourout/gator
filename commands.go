package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exist := c.handlers[cmd.name]
	if !exist {
		return fmt.Errorf("command does not exist: %s", cmd.name)
	}

	err := handler(s, cmd)
	return err
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.handlers[name] = f
}
