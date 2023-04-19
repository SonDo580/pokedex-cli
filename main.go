package main

import (
	"fmt"
)

func main() {
	stop := false

	for ; !stop ; {
		fmt.Printf("Pokedex > ")

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			stop = true
		}
	}
}