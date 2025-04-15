package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *Config, parameter string) error {

	url := "https://pokeapi.co/api/v2/location-area" //default API url
	if cfg.Next != "" {
		url = cfg.Next //if the next URL exists, set it
	}

	resp, err := cfg.Client.GetRequest(url) //make request
	if err != nil {
		return err
	}
	//defer now happens in getrequest method

	var locationsresp LocationAreaResponse //initialize a variable struct to hold the json responses in

	err = json.Unmarshal(resp, &locationsresp) // unmarshals the response into locationsresponse form
	if err != nil {
		return err
	}
	cfg.Next = locationsresp.Next //update the URLs of the config struct
	cfg.Previous = locationsresp.Previous

	for _, result := range locationsresp.Results { //print da reesults
		fmt.Println("---", result.Name, "---")
	}
	return nil
}
