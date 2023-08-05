package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		err := command.callback()
		if err != nil {
			return
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
	callback    func() error
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
			description: "Lists some location areas",
			callback:    callbackMap,
		},
	}
}
