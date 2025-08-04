package cmd

import (
	"errors"
	"fmt"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CommandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}

	name := args[0]
	location, err := cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
