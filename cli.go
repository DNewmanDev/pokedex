package main

import "strings"

func cleanInput(input string) []string {
	if len(input) == 0 {
		return nil
	}
	trimmedInput := strings.Fields(strings.ToLower(input))

	return trimmedInput
}

type commandCallback func(cfg *Config, parameter string) error
type cliCommand struct {
	name        string
	description string
	callback    commandCallback
}

var pokecommands map[string]cliCommand

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
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores an area for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
