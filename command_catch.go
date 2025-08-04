package main

import (
	"errors"
	"fmt"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid arguments")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	// TODO: Implement catch logic
	return nil
}
