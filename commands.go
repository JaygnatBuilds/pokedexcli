package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type Pokemon struct {
	name     string
	base_exp int
}

var pokedex = map[string]Pokemon{}

func commandExit(cfg *config, param string) error {

	// Parameter check
	if param != "" {
		fmt.Println("exit command accepts no parameter")
		return nil
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param string) error {

	// Parameter check
	if param != "" {
		fmt.Println("help command accepts no parameter")
		return nil
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands")
	fmt.Println("--------")
	for _, command := range getCommands() {
		fmt.Printf("%s : %s\n", command.name, command.description)
	}
	fmt.Println("--------")
	return nil
}

func commandMap(cfg *config, param string) error {

	// Parameter check
	if param != "" {
		fmt.Println("map command accepts no parameter")
		return nil
	}

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

func commandMapb(cfg *config, param string) error {

	// Parameter check
	if param != "" {
		fmt.Println("mapb command accepts no parameter")
		return nil
	}

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

func commandClear(cfg *config, param string) error {

	// Parameter check
	if param != "" {
		fmt.Println("clear command accepts no parameter")
		return nil
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}

func commandExplore(cfg *config, area string) error {

	data, err := cfg.client.ListPokemonEncounters(area)

	if err != nil {
		return err
	}

	fmt.Printf("Listing pokemon in %v...\n", data.Location.Name)

	for _, result := range data.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}

func commandCatchPokemon(cfg *config, name string) error {

	data, err := cfg.client.CatchPokemon(name)

	if err != nil {
		return err
	}

	// check if pokemon is already in pokedex
	if _, ok := pokedex[data.Name]; ok {
		fmt.Printf("%v is already in your pokedex.\n", data.Name)
		return nil
	} else {
		fmt.Printf("Throwing Pokeball at %v...\n", data.Name)
		fmt.Printf("%v base experience is %v\n", data.Name, data.Base_Exp)

		exp_normalize := data.Base_Exp / 10
		odds_final := 100 - exp_normalize

		fmt.Printf("exp_normalization: %d\n", exp_normalize)
		fmt.Printf("final odds: %d\n", odds_final)

		catch_roll := rand.Intn(100)
		fmt.Printf("catch roll : %d\n", catch_roll)
		if catch_roll <= odds_final {
			fmt.Println("Caught!")
			pokedex[data.Name] = Pokemon{data.Name, data.Base_Exp}
		} else {
			fmt.Println("Miss!")
		}
	}

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
		"explore": {
			name:        "explore",
			description: "list all pokemon in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    commandCatchPokemon,
		},
	}
}
