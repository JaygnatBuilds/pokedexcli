package main

import (
	"time"

	"github.com/JaygnatBuilds/pokedexcli/internal/pokeapi"
)

type config struct {
	next   string
	prev   string
	client *pokeapi.Client
}

func main() {

	// initialize config client
	cfg := &config{
		client: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	startRepl(cfg)
}
