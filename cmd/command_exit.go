package cmd

import (
	"fmt"
	"os"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CommandExit(cfg *config.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
