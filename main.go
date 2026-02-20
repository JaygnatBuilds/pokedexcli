package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initialize REPL loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt text
		fmt.Print("Pokedex > ")

		// Scan for commands and capture first string in input
		scanner.Scan()
		text := CleanInput(scanner.Text())
		command := text[0]
		fmt.Printf("Your command was: %s\n", command)

	}
}
