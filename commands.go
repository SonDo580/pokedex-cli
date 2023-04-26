package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SonDo580/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	description string
	callback    func()
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

func callbackHelp() {
	availableCommands := getAvailableCommands()
	
	for commandName, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", commandName, command.description)
	}

	fmt.Println("----------")
}

func callBackExit() {
	os.Exit(0)
}

func callBackMap() {
	pokeAPIClient := pokeapi.NewClient()

	data, err :=  pokeAPIClient.GetLocationsData()
	if err != nil {
		log.Fatal(err)
	}

	printLocations(data)
}

func printLocations(data pokeapi.LocationsResponse) {
	fmt.Println("Locations:")
	for _, location := range data.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	fmt.Println("----------")
}

func callbackMapBack() {
}