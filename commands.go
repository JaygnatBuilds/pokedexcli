package main

import (
	"fmt"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(cfg *config) error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands")
	fmt.Println("--------")
	for _, command := range getCommands() {
		fmt.Printf("%s : %s\n", command.name, command.description)
	}
	fmt.Println("--------")
	return nil
}

func commandMap(cfg *config) error {

	data, err := cfg.client.ListLocationAreas(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = data.Next
	cfg.prev = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {

	if cfg.prev == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	data, err := cfg.client.ListLocationAreas(cfg.prev)
	if err != nil {
		return err
	}

	cfg.next = data.Next
	cfg.prev = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandClear(cfg *config) error {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"clear": {
			name:        "clear",
			description: "clear console screen",
			callback:    commandClear,
		},
	}
}
