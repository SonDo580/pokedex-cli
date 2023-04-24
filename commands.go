package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getAvailableCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callBackExit,
		},
	}
}

func callbackHelp() {
	availableCommands := getAvailableCommands()
	
	for _, command := range(availableCommands) {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	fmt.Println("----------")
}

func callBackExit() {

}