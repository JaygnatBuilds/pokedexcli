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

	// read response body into byte[] data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// add result to cache and return result
	c.cache.Add(url, data)
	return locationsResp, nil

}
