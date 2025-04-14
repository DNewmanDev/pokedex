package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleaned := cleanInput(input)

		if cmd, exists := pokecommands[cleaned[0]]; exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknown command")
		}

	}

}

func cleanInput(input string) []string {
	if len(input) == 0 {
		return nil
	}
	trimmedInput := strings.Fields(strings.ToLower(input))

	return trimmedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var pokecommands map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range pokecommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}

func init() {

	pokecommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
