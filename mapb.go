package main

import (
	"encoding/json"
	"fmt"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := cfg.Client.GetRequest(cfg.Previous) //make request
	if err != nil {
		return err
	}
	//defer now happens in the get request

	var locationsresp LocationAreaResponse     //init json response container
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
