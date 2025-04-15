package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, parameter string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + parameter

	resp, err := cfg.Client.GetRequest(url)
	if err != nil {
		return err
	}
	var pokeinfo PokemonInfo

	err = json.Unmarshal(resp, &pokeinfo)
	if err != nil {
		return err
	}
	randomseed := rand.Intn(pokeinfo.BaseExperience)
	fmt.Println("Throwing a Pokeball at " + pokeinfo.Name + " ...")
	fmt.Println(randomseed)
	fmt.Println(pokeinfo.BaseExperience / 2)
	if randomseed > pokeinfo.BaseExperience/2+pokeinfo.BaseExperience/6 {
		print(pokeinfo.Name + " was caught!\n")
		//add to map
	} else {
		fmt.Println(pokeinfo.Name + " got away!")
	}
	return nil
}
