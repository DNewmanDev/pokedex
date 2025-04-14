package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *Config) error {

	url := "https://pokeapi.co/api/v2/location-area" //default API url
	if cfg.Next != "" {
		url = cfg.Next //if the next URL exists, set it
	}

	resp, err := http.Get(url) //make request
	if err != nil {
		return err
	}
	defer resp.Body.Close() // make sure the request is closed at finishing

	var locationsresp LocationAreaResponse                                    //initialize a variable struct to hold the json responses in
	if err := json.NewDecoder(resp.Body).Decode(&locationsresp); err != nil { //decode the json and throw any errors if they're there
		return err
	}
	cfg.Next = locationsresp.Next
	cfg.Previous = locationsresp.Previous

	for _, result := range locationsresp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
