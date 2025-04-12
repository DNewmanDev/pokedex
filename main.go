package main

import (
	"strings"
)

func main() {

}

func cleanInput(input string) []string {
	if len(input) == 0 {
		return nil
	}
	trimmedInput := strings.Fields(strings.ToLower(input))

	return trimmedInput
}
