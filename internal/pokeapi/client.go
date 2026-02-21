package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client
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
	}
}

func (c *Client) ListLocationAreas(url string) (LocationAreaResponse, error) {

	// if no next or prev urls are stored in config, use base location-area url
	if url == "" {
		url = c.BaseURL + "/location-area"
	}

	// make get API call
	res, err := c.httpClient.Get(url)

	// error check API call
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	defer res.Body.Close()

	// decode response data into LocationAreaResponse object
	var data LocationAreaResponse
	json.NewDecoder(res.Body).Decode(&data)

	return data, nil
}
