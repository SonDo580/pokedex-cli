package main

import (
	"time"

	"github.com/SonDo580/pokedex-cli/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func main() {
	appConfig := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&appConfig)
}
