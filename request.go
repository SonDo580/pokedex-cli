package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getLocationsData() (locationsResponse, error) {
	endpoint := "/location"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationsResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationsResponse{}, fmt.Errorf("%v: %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationsResponse{}, err
	}

	locationsData := locationsResponse{}
	err = json.Unmarshal(data, &locationsData)
	if err != nil {
		return locationsResponse{}, err
	}

	return locationsData, nil
}