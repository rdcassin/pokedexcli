package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetAreaInfo(areaIDOrName string) (AreaInfoAPIResults, error) {
	url := baseURL + "location-area/" + areaIDOrName + "/"

	if cachedData, exists := c.cache.Get(url); exists {
		results := AreaInfoAPIResults{}
		err := json.Unmarshal(cachedData, &results)
		if err != nil {
			return AreaInfoAPIResults{}, err
		}
		return results, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaInfoAPIResults{}, err
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaInfoAPIResults{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaInfoAPIResults{}, err
	}

	c.cache.Add(url, bodyBytes)

	results := AreaInfoAPIResults{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return AreaInfoAPIResults{}, err
	}

	return results, nil
}