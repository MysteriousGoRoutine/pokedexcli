package config

import "github.com/MysteriousGoRoutine/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiClient    pokeapi.Client
	CaughtPokemon    map[string]pokeapi.Pokemon
	NextLocationsURL *string
	PrevLocationsURL *string
}
