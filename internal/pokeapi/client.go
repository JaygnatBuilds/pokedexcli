package pokeapi

import (
	"encoding/json"
	"fmt"
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

	var data LocationAreaResponse

	// if no next or prev urls are stored in config, use base location-area url
	if url == "" {
		url = c.BaseURL + "/location-area"
	}

	// make get API call
	res, err := c.httpClient.Get(url)

	// error check API call
	if err != nil {
		return LocationAreaResponse{}, err
	}
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// decode response data into LocationAreaResponse object
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return LocationAreaResponse{}, err
	}

	return data, nil
}
