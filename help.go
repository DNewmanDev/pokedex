package main

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range pokecommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}
