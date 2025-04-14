package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	cfg := &Config{}                      //initializing the blank configuration
	scanner := bufio.NewScanner(os.Stdin) //object that listens for input

	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()          //wait for input
		input := scanner.Text() //register input
		if len(input) == 0 {    //if they send nothing, do nothing
			continue
		}
		cleaned := cleanInput(input) //Clean the input

		command, ok := pokecommands[cleaned[0]] //pull the command
		if !ok {                                //iff the command isn't in the list, print not found
			fmt.Println("Command not found")
			continue

		}
		err := command.callback(cfg) //check the callback function and for error
		if err != nil {
			fmt.Println("Error: ", err)
		}

	}

}

func cleanInput(input string) []string {
	if len(input) == 0 {
		return nil
	}
	trimmedInput := strings.Fields(strings.ToLower(input))

	return trimmedInput
}

type commandCallback func(cfg *Config) error
type cliCommand struct {
	name        string
	description string
	callback    commandCallback
}
type Config struct {
	Next     string
	Previous string
}
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var pokecommands map[string]cliCommand

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range pokecommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}
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
func init() {

	pokecommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
	}
}
