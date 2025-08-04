package cmd

import (
	"strings"

	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"explore": {
			Name:        "explore <location name>",
			Description: "Explore a location",
			Callback:    CommandExplore,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "View details about a caught Pokemon",
			Callback:    CommandInspect,
		},
		"catch": {
			Name:        "catch <pokemon name>",
			Description: "Catch a pokemon",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "prints a list of all the names of the Pokemon the user has caught",
			Callback:    CommandPokedex,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}
