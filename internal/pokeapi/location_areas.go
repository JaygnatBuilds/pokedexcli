// Functions related to the location-areas pokeAPI endpoint
// ListLocationAreas : returns a list of location areas
// ListPokemonEncounters : returns a list of pokemon a location-area has

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocationAreas(url string) (LocationAreaResponse, error) {

	// if no next or prev urls are stored in config, use base location-area url
	if url == "" {
		url = c.BaseURL + "/location-area"
	}

	// Check if api call response value is already in cache
	if value, ok := c.cache.Get(url); ok {

		locationsResp := LocationAreaResponse{}
		err := json.Unmarshal(value, &locationsResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationsResp, nil
	}

	// GET API call
	res, err := c.httpClient.Get(url)

	// error check API call
	if err != nil {
		return LocationAreaResponse{}, err
	}
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// read response body into byte[] data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// convert byte[] data to LocationAreaResponse object
	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// add result to cache and return result
	c.cache.Add(url, data)
	return locationsResp, nil

}

func (c *Client) ListPokemonEncounters(area string) (PokemonEncounterResponse, error) {

	// return error if no location parameter is provided
	if area == "" {
		return PokemonEncounterResponse{}, fmt.Errorf("location parameter blank. Please provide location name or ID.")
	}

	url := c.BaseURL + "/location-areas/" + area

	// Check if api call response value is already in cache
	if value, ok := c.cache.Get(url); ok {

		pokemonResponse := PokemonEncounterResponse{}
		err := json.Unmarshal(value, &pokemonResponse)
		if err != nil {
			return PokemonEncounterResponse{}, err
		}

		return pokemonResponse, nil
	}

	// GET API call
	res, err := c.httpClient.Get(url)

	// error check API call
	if err != nil {
		return PokemonEncounterResponse{}, err
	}
	if res.StatusCode > 299 {
		return PokemonEncounterResponse{}, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// read response body into byte[] data
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonEncounterResponse{}, err
	}

	// convert byte[] data to PokemonEncounterResponse object
	pokemonResponse := PokemonEncounterResponse{}
	err = json.Unmarshal(data, &pokemonResponse)

	if err != nil {
		return PokemonEncounterResponse{}, err
	}

	// add result to cache and return result
	c.cache.Add(url, data)
	return pokemonResponse, nil

}
