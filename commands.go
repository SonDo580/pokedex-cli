package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SonDo580/pokedex-cli/pokeapi"
)

type cliCommand struct {
	description string
	callback    func(*config)
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

func callbackHelp(appConfig *config) {
	availableCommands := getAvailableCommands()
	
	for commandName, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", commandName, command.description)
	}

	fmt.Println("----------")
}

func callBackExit(appConfig *config) {
	os.Exit(0)
}

func callBackMap(appConfig *config) {
	pokeAPIClient := pokeapi.NewClient()

	data, err :=  pokeAPIClient.GetLocationsData()
	if err != nil {
		log.Fatal(err)
	}

	printLocations(data)
}

func callbackMapBack(appConfig *config) {
}

func printLocations(data pokeapi.LocationsResponse) {
	fmt.Println("Locations:")
	for _, location := range data.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	fmt.Println("----------")
}