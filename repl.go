package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(appConfig *config) {
	scanner := bufio.NewScanner(os.Stdin)

	availableCommands := getAvailableCommands()

	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		command.callback(appConfig)
	}
}
 
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}