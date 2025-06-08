package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetMap(pageURL *string) (LocationArea, error) {
	url := baseURL + "location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, exists := c.cache.Get(url); exists {
		results := LocationArea{}
		err := json.Unmarshal(cachedData, &results)
		if err != nil {
			return LocationArea{}, err
		}
		
		return results, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, errors.New("this area does not exist or is not available at this time")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, bodyBytes)

	results := LocationArea{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return LocationArea{}, err
	}

	return results, nil
}