package main

import (
	"fmt"
	"log"
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
		"map": {
			description: "Displays the next 20 locations",
			callback: callBackMap,
		},
		"mapb": {
			description: "Displays the previous 20 locations",
			callback: callbackMapBack,
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

func callBackMap() error {
	pokeAPIClient := NewClient()

	data, err :=  pokeAPIClient.getLocationsData()
	if err != nil {
		log.Fatal(err)
	}

	printLocations(data)

	return nil
}

func printLocations(data locationsResponse) {
	fmt.Println("Locations:")
	for _, location := range data.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	fmt.Println("----------")
}

func callbackMapBack() error {
	return nil
}