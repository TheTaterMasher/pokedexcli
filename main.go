package main

import (
	"time"

	"github.com/TheTaterMasher/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
	catchThreshold          int
}

func main() {
	cfg := config{
		pokeapiClient:  pokeapi.NewClient(time.Hour),
		caughtPokemon:  make(map[string]pokeapi.Pokemon),
		catchThreshold: 25,
	}
	startRepl(&cfg)
}
