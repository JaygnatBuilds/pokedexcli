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

func NewClient() *Client {

	return &Client{
		BaseURL:    "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(5 * time.Second),
	}
}
