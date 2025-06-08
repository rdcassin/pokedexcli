package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonIDOrName string) (PokemonInfoAPIResults, error) {
	url := baseURL + "pokemon/" + pokemonIDOrName + "/"

	if cachedData, exists := c.cache.Get(url); exists {
		results := PokemonInfoAPIResults{}
		err := json.Unmarshal(cachedData, &results)
		if err != nil {
			return PokemonInfoAPIResults{}, err
		}

		return results, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfoAPIResults{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfoAPIResults{}, err
	}
	defer resp.Body.Close()
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfoAPIResults{}, err
	}

	c.cache.Add(url, bodyBytes)

	results := PokemonInfoAPIResults{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return PokemonInfoAPIResults{}, err
	}

	return results, nil
}