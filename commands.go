package main

import (
	"fmt"
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
			callback:    callBackMap,
		},
		"mapb": {
			description: "Displays the previous 20 locations",
			callback:    callbackMapBack,
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
	data, err := appConfig.pokeapiClient.GetLocationsData(appConfig.nextLocationURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	printLocations(data)
	appConfig.nextLocationURL = data.Next
	appConfig.prevLocationURL = data.Previous
}

func callbackMapBack(appConfig *config) {
	if appConfig.prevLocationURL == nil {
		fmt.Println("You're on the first page")
		return
	}

	data, err := appConfig.pokeapiClient.GetLocationsData(appConfig.prevLocationURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	printLocations(data)
	appConfig.nextLocationURL = data.Next
	appConfig.prevLocationURL = data.Previous
}

func printLocations(data pokeapi.LocationsResponse) {
	fmt.Println("Locations:")
	for _, location := range data.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	fmt.Println("----------")
}
