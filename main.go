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
		Client: client,
	} //initializing the blank configuration
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
