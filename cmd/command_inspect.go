package cmd

import (
	"errors"
	"fmt"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CommandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.CaughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, type_ := range pokemon.Types {
		fmt.Printf("  -%s\n", type_.Type.Name)
	}
	fmt.Println()

	return nil
}
