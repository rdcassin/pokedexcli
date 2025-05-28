package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client)GetMap(pageURL *string) (LocationAreaAPIResults, error) {
	url := baseURL + "location-area/"

	if pageURL != nil {
		url = *pageURL
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

	results := LocationAreaAPIResults{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&results)
	if err != nil {
		return LocationAreaAPIResults{}, err
	}

	return results, nil
}