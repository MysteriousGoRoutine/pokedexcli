package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MysteriousGoRoutine/pokedexcli/cmd"
	"github.com/MysteriousGoRoutine/pokedexcli/internal/config"
)

func startRepl(cfg *config.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cmd.CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cmd.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
