package main

import (
	"fmt"
	"os"
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
	for commandName, command := range commands {
		fmt.Printf("%s : %s\n", commandName, command.description)
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

// TODO : Clear function that clears console
