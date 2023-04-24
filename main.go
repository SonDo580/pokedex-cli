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
		case "exit":
			os.Exit(0)
		}

		fmt.Println("Echoing: ", cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}