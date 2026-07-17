// Functions related to the pokemon pokeAPI endpoint
// ListLocationAreas : returns a list of location areas

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) CatchPokemon(name string) (PokemonInfoResponse, error) {

	if name == "" {
		return PokemonInfoResponse{}, fmt.Errorf("Pokemon parameter blank. Please provide Pokemon name or ID.")
	}

	url := c.BaseURL + "/pokemon/" + name

	// Check if api call response value is already in cache
	if value, ok := c.cache.Get(url); ok {

		pokemonInfoResponse := PokemonInfoResponse{}
		err := json.Unmarshal(value, &pokemonInfoResponse)
		if err != nil {
			return PokemonInfoResponse{}, err
		}

		return pokemonInfoResponse, nil
	}

	// GET API call
	res, err := c.httpClient.Get(url)

	// error check API call
	if err != nil {
		return PokemonInfoResponse{}, err
	}
	if res.StatusCode > 299 {
		return PokemonInfoResponse{}, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// read response body into byte[] data
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonInfoResponse{}, err
	}

	// convert byte[] data to PokemonEncounterResponse object
	pokemonInfoResponse := PokemonInfoResponse{}
	err = json.Unmarshal(data, &pokemonInfoResponse)

	if err != nil {
		return PokemonInfoResponse{}, err
	}

	// add result to cache and return result
	c.cache.Add(url, data)
	return pokemonInfoResponse, nil

}
