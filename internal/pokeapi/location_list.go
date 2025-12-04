package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowLocations{}, err
	}

	locations := ShallowLocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return ShallowLocations{}, err
	}

	return locations, nil
}
