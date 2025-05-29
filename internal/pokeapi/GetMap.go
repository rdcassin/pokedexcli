package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetMap(pageURL *string) (LocationAreaAPIResults, error) {
	url := baseURL + "location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.cache.Get(url); found {
		results := LocationAreaAPIResults{}
		err := json.Unmarshal(cachedData, &results)
		if err != nil {
			return LocationAreaAPIResults{}, err
		}
		return results, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaAPIResults{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaAPIResults{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaAPIResults{}, err
	}

	c.cache.Add(url, bodyBytes)

	results := LocationAreaAPIResults{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return LocationAreaAPIResults{}, err
	}

	return results, nil
}