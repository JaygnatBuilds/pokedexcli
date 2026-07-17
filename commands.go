// REPL command function definitions
// exit : exit the REPL
// help : print available commands and usage
// map : display next 20 map locations
// mapb : display previous 20 map locations
// clear : clear console screen
// explore : list all pokemon in an area
// catch : attempt to catch a pokemon
// inspect : inspect a pokemons attributes
// pokedex : list all pokemon in pokedex

package main

import (
	"encoding/json"
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
	Name   string
	Height int
	Weight int
	Stats  []PokemonStat
	Types  []PokemonType
}

type PokemonStat struct {
	Name      string
	Base_stat int
}

type PokemonType struct {
	Name string
}

// map for storing pokemon
var pokedex = map[string]Pokemon{}

// exit command : exit the REPL
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

// help command : list available commands and usage
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

// map command : display next 20 map locations from API
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

// mapb command : display previous 20 map locations from API
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

// clear command : clear REPL console screen
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

// explore command : list all pokemon in an area
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

// catch command : attempt to catch a pokemon
func commandCatchPokemon(cfg *config, name string) error {

	// call catch pokemon API function
	data, err := cfg.client.CatchPokemon(name)

	if err != nil {
		return err
	}

	// check if pokemon is already in pokedex
	if _, ok := pokedex[data.Name]; ok {
		fmt.Printf("%v is already in your pokedex.\n", data.Name)
		return nil
	} else {
		// Simulate catching the pokemon with chances tied to pokemon base exp
		fmt.Printf("Throwing a Pokeball at %v...\n", data.Name)
		fmt.Printf("%v base experience is %v\n", data.Name, data.Base_Exp)

		exp_normalize := data.Base_Exp / 10
		odds_final := 100 - exp_normalize

		fmt.Printf("exp_normalization: %d\n", exp_normalize)
		fmt.Printf("final odds: %d\n", odds_final)

		catch_roll := rand.Intn(100)
		fmt.Printf("catch roll : %d\n", catch_roll)

		// If catch is successful, add pokemon to pokedex
		if catch_roll <= odds_final {

			fmt.Println("Caught! Adding to PokeDex...")
			pokemon := Pokemon{}
			pokemon.Name = data.Name
			pokemon.Height = data.Height
			pokemon.Weight = data.Weight
			for _, s := range data.Stats {
				stat := PokemonStat{s.Stat.Name, s.BaseStat}
				pokemon.Stats = append(pokemon.Stats, stat)
			}
			for _, s := range data.Types {
				pokeType := PokemonType{s.Type.Name}
				pokemon.Types = append(pokemon.Types, pokeType)
			}

			pokedex[data.Name] = pokemon

		} else {
			fmt.Println("Miss!")
		}
	}

	return nil
}

// inspect command : inspect a pokemons attributes stored in pokedex
func commandInspectPokemon(cfg *config, name string) error {

	// parameter check
	if name == "" {
		fmt.Println("Name parameter empty, please provide pokemon name with this command.")
		return nil
	}

	// check if pokemon is in pokedex
	if _, ok := pokedex[name]; ok {
		// convert pokemon object to json bytes
		data, err := json.MarshalIndent(pokedex[name], "", "	")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	} else {
		fmt.Println("you have not caught that pokemon...")
		return nil
	}
	return nil
}

// pokedex command : list all pokemon in pokedex
func commandInspectPokedex(cfg *config, param string) error {

	// parameter check
	if param != "" {
		fmt.Println("Inspect Pokedex command accepts no parameter")
		return nil
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
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemons attributes",
			callback:    commandInspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list pokemon in pokedex",
			callback:    commandInspectPokedex,
		},
	}
}
