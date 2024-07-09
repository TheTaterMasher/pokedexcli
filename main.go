package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		if input == "" {
			continue
		}

		cliCommands := getCliCommands()
		command := cliCommands[input]

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
	return nil
}

func commandExit() error {
	return errors.New("exiting pokedex")
}
