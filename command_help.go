package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	cliCommands := getCliCommands()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}
