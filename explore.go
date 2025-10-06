package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(cfg *Config, parameter string) error {


	url := "https://pokeapi.co/api/v2/location-area/" + parameter

	resp, err := cfg.Client.GetRequest(url)
	if err != nil {
		return err
	}

	var exploreresp ExploreResponse

	err = json.Unmarshal(resp, &exploreresp)
	if err != nil {
		return err
	}
	fmt.Println("You found:")
	for _, result := range exploreresp.PokemonEncounters {
		fmt.Println("-" + result.Pokemon.Name)
	}

	return nil
}
