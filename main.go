package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Printf("Your command was: %s \n", cleaned[0])
	}
}

func cleanInput(input string) []string {
	if len(input) == 0 {
		return nil
	}
	trimmedInput := strings.Fields(strings.ToLower(input))

	return trimmedInput
}
