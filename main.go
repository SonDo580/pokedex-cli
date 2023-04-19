package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stop := false
	scanner := bufio.NewScanner(os.Stdin)

	for ; !stop ; {
		fmt.Printf("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			stop = true
		}
	}
}