package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) <= 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommand := getCommands()

		command, ok := availableCommand[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type command struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Prints the help menu with available commands",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists some NEXT location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists some PREV location areas",
			callback:    callbackMapB,
		},
	}
}
