package main

import (
	"fmt"
	"github.com/AmelAbema/Pokedex/internal/pokeapi"
	"log"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	areas, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas: ")
	for _, area := range areas.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	return nil
}
