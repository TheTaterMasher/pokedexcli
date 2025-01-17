package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %v\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
