package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	areas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas: ")
	for _, area := range areas.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = areas.Next
	cfg.prevLocationAreaURL = areas.Next
	return nil
}
func callbackMapB(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("U are on the first page")
	}
	areas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas: ")
	for _, area := range areas.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = areas.Next
	cfg.prevLocationAreaURL = areas.Next
	return nil
}
