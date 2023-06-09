package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationsData(pageURL *string) (LocationsResponse, error) {
	endpoint := "/location"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check cache
	cachedData, ok := c.cache.Get(fullURL)
	if ok {
		locationsData := LocationsResponse{}
		err := json.Unmarshal(cachedData, &locationsData)
		if err != nil {
			return LocationsResponse{}, err
		}

		return locationsData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationsResponse{}, fmt.Errorf("%v: %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	locationsData := LocationsResponse{}
	err = json.Unmarshal(data, &locationsData)
	if err != nil {
		return LocationsResponse{}, err
	}

	// Cache success data
	c.cache.Add(fullURL, data)

	return locationsData, nil
}
