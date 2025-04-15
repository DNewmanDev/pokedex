package main

import "fmt"

func commandList(cfg *Config, parameter string) error {
	fmt.Println("Your pokedex: ")

	for _, result := range cfg.Pokedex.pokedex {
		fmt.Println(" - " + cfg.Pokedex.pokedex[result.Name].Name)

	}

	// for _, result := cfg.Pokedex.pokedex{
	// 	fmt.Println(result.)
	// }
	return nil
}
