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
		fmt.Println("You Typed >>> ", text)
	}
}

func cleanInput(str string) []string {

	lowered := strings.ToLower(str)

	words:= strings.Fields(lowered)

	return words
}
