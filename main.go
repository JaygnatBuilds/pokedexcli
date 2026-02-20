package main

import (
	"bufio"
	"fmt"
	"os"
)

var commands map[string]cliCommand

func main() {

	// initialize commands map
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print available commands",
			callback:    commandHelp,
		},
	}

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
		function, ok := commands[command]

		// if command exists in commands map, run function callback, else throw error message
		if ok {
			function.callback()
		} else {
			fmt.Println("Please enter valid command (\"Help\" is a valid command)")
		}

	}
}
