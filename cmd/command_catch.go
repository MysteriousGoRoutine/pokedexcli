package cmd

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CommandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}
