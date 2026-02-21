package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JaygnatBuilds/pokedexcli/internal/pokeapi"
)

var commands map[string]cliCommand

type config struct {
	next   string
	prev   string
	client *pokeapi.Client
}

func main() {

	// initialize config client
	cfg := &config{
		client: pokeapi.NewClient(),
	}

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
		"map": {
			name:        "map",
			description: "display next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display previous 20 map locations",
			callback:    commandMapb,
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
			function.callback(cfg)
		} else {
			fmt.Println("Please enter valid command (\"Help\" is a valid command)")
		}

	}
}
