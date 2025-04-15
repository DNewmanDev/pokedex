package main

import "fmt"

func commandList(cfg *Config, parameter string) error {
	fmt.Println("Your pokedex: ")

	for _, result := range cfg.Pokedex.pokedex {
		fmt.Println(" - " + cfg.Pokedex.pokedex[result.Name].Name)

	}

	return nil
}
