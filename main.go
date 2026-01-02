package main

import (
	"fmt"
	"log"

	"github.com/dvantourout/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error when reading config file: %v", err)
	}

	fmt.Printf("config file:\n%v\n", cfg)

	if err := cfg.SetUser("dvantourout"); err != nil {
		log.Fatalf("error when setting user name: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error when reading config file: %v", err)
	}

	fmt.Printf("config file:\n%v\n", cfg)
}
