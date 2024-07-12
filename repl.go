package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()

		cleanedinput := cleanInput(input)

		if len(cleanedinput) == 0 {
			continue
		}

		command := cliCommands[cleanedinput[0]]

		if command.name == "" {
			fmt.Println("invalid command, for a list of commands type help")
			continue
		}

		err := command.callback()

		if err != nil {
			fmt.Println("Exiting Pokedex")
			break
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower((str))
	cleaned := strings.Fields(lowered)
	return cleaned
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next map areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous map areas",
			callback:    commandMapB,
		},
	}
}

func commandHelp() error {
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

func commandExit() error {
	return errors.New("exiting pokedex")
}

func commandMap() error {
	return nil
}

func commandMapB() error {
	return nil
}
