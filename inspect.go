package main

import (
	"fmt"
	"strconv"
)

func commandInspect(cfg *Config, parameter string) error {

	val, ok := cfg.Pokedex.Get(parameter)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name: " + val.Name)
	fmt.Println("Height: " + strconv.Itoa(val.Height))
	fmt.Println("Weight: " + strconv.Itoa(val.Weight))
	fmt.Println("Stats: ")
	for _, result := range val.Stats {
		fmt.Println(" -" + result.Stat.Name + ": " + strconv.Itoa(result.BaseStat))
	}
	fmt.Println("Types: ")
	for _, result := range val.Types {
		fmt.Println(" - " + result.Type.Name)
	}

	//fmt.Println(val.Types)
	//get the pokemon from the pokedex
	//if not there, you have not caught htis pokemon
	//print name height weight stats types

	return nil
}
