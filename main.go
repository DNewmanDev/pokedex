package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	client := NewClient(5 * time.Minute) //creates client for HTTP requests, lasts 5 minutes
	cfg := &Config{
		Client:  client,
		Pokedex: NewPokedex(),
	} //initializing the blank configuration

	scanner := bufio.NewScanner(os.Stdin) //object that listens for input

	// Check the Pokedex for caught Pokémon: THIS IS FOR THE LIST FUNCTIONcatch gya
	// fmt.Println("Caught Pokémon in Pokedex:")
	// for name, _ := range cfg.Pokedex.pokedex {
	// 	// fmt.Printf("%s: %+v\n", name, info)
	// 	fmt.Println(name)
	// }
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()          //wait for input
		input := scanner.Text() //register input
		parameter := ""
		if len(input) == 0 { //if they send nothing, do nothing
			continue
		}
		cleaned := cleanInput(input) //Clean the input
		if len(cleaned) > 1 {
			parameter = cleaned[1]
		}
		command, ok := pokecommands[cleaned[0]] //pull the command
		if !ok {                                //iff the command isn't in the list, print not found
			fmt.Println("Command not found")
			continue
		}
		//if command is explore pull cleaned [0] AND cleaned[1]

		err := command.callback(cfg, parameter) //check the callback function and for error
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
