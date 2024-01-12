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

		fmt.Print("Your Input > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Printf("Unknown command: %v\n", commandName)
			fmt.Printf("To see the list of available commands, type 'help'\n")
			continue
		}

		err := command.callback(cfg, args...)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the available commands",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Prints a page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {

	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
