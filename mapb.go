package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var locationsresp LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationsresp); err != nil {
		return err
	}

	cfg.Next = locationsresp.Next
	cfg.Previous = locationsresp.Previous

	for _, result := range locationsresp.Results {
		fmt.Println(result.Name)
	}

	return nil

}
