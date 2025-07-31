package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help information",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// Clean the input string
		commands := cleanInput(scanner.Text())
		if len(commands) == 0 {
			continue
		}
		switch commands[0] {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		default:
			fmt.Println("Unknown command:", commands[0])
		}
		// fmt.Printf("Your command was: %s\n", words[0])

	}
}
