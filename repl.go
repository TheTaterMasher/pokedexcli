package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		command, ok := cliCommands[cleanedinput[0]]

		if !ok {
			fmt.Println("invalid command, for a list of commands type 'help'")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower((str))
	cleaned := strings.Fields(lowered)
	return cleaned
}
