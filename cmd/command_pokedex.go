package cmd

import (
	"errors"
	"fmt"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CommandPokedex(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		return errors.New("pokedex command takes no arguments")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
