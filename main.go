package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dvantourout/gator/internal/config"
	"github.com/dvantourout/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error when reading config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	dbQueries := database.New(db)

	s := &state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)

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
