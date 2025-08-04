package main

import (
	"time"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
	"github.com/MysteriousGoRoutine/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
