package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands")
	fmt.Println("--------")
	for commandName, command := range commands {
		fmt.Printf("%s : %s\n", commandName, command.description)
	}
	fmt.Println("--------")
	return nil
}

// TODO : Clear function that clears console
