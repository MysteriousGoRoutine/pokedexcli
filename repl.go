package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// Clean the input string
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
