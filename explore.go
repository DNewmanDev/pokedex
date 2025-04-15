package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(cfg *Config, parameter string) error {
	//somehow the user input has to get here for
	//us to know where to explore
	//for now while building exp functionality
	//will use dummy url

	url := "https://pokeapi.co/api/v2/location-area/" + parameter
	//check later to see if I need to add
	//extra fields to config

	resp, err := cfg.Client.GetRequest(url)
	if err != nil {
		return err
	}

	var exploreresp ExploreResponse

	err = json.Unmarshal(resp, &exploreresp)
	if err != nil {
		return err
	}
	fmt.Println("Here's ya pokemon")
	for _, result := range exploreresp.PokemonEncounters {
		fmt.Println("-" + result.Pokemon.Name)
	}

	return nil
}
