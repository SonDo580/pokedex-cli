package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	description string
	callback    func() error
}

func getAvailableCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			description: "Displays the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			description: "Exit the Pokedex",
			callback:    callBackExit,
		},
	}
}

func callbackHelp() error {
	availableCommands := getAvailableCommands()
	
	for commandName, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", commandName, command.description)
	}

	fmt.Println("----------")

	return nil
}

func callBackExit() error {
	os.Exit(0)
	return nil
}