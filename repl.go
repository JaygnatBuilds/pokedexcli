package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {

	wordsLower := strings.ToLower(text)
	words := strings.Fields(wordsLower)

	return words
}

func startRepl(cfg *config) {

	// initialize REPL loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt text
		fmt.Print("Pokedex > ")

		// Scan for commands and capture first string in input
		scanner.Scan()
		text := CleanInput(scanner.Text())
		command := text[0]

		// Obtain function callback from commands map
		function, ok := getCommands()[command]

		// if command exists in commands map, run function callback, else throw error message
		if ok {

			err := function.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {

			fmt.Println("Please enter valid command (\"Help\" is a valid command)")
			continue

		}

	}
}
