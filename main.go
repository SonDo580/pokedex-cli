package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("Welcome to the Pokedex help menu!")
			fmt.Println("Here are the available commands:")
			fmt.Println("- help")
			fmt.Println("- exit")
			fmt.Println("----------")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}