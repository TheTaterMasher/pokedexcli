package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("your on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %vs\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
