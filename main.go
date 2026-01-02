package main

import (
	"log"
	"os"

	"github.com/dvantourout/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error when reading config file: %v", err)
	}

	s := &state{
		config: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", loginHandler)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("need at least one argument")
	}

	cmd := command{
		name: args[1],
	}
	if len(args) > 2 {
		cmd.args = args[2:]
	}

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
