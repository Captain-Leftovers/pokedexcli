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

		fmt.Print("Your Input > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Printf("Unknown command: %v\n", commandName)
			fmt.Printf("To see the list of available commands, type 'help'\n")
			continue
		}

		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the available commands",
			callback:    callbackHelp,
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
