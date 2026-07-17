// PokeAPI client driver file
// API client struct as well as JSON response structs
// NewClient : function to initialize a new pokeAPI client
package pokeapi

import (
	"net/http"
	"time"

	"github.com/JaygnatBuilds/pokedexcli/internal/pokecache"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonEncounterResponse struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonInfoResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Base_Exp int    `json:"base_experience"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	Stats    []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// NewClient : initialize new API client
func NewClient(timeout, cacheInterval time.Duration) *Client {

	return &Client{
		BaseURL: "https://pokeapi.co/api/v2",
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
