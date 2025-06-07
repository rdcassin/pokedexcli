package main

import (
	"fmt"
)

func commandExplore(c *config, areaIDOrName string) error {
	areaInfo, err := c.pokeapiClient.GetAreaInfo(areaIDOrName)
	if err != nil {
		return err
	}

	possiblePokemon := areaInfo.PokemonEncounters

	if len(possiblePokemon) == 0 {
		fmt.Printf("No Pokemon found in area: %s\n", areaIDOrName)
		return nil
	}

	fmt.Printf("Exploring %s...\n", areaIDOrName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range possiblePokemon {
		fmt.Printf("- %s\n", pokemon.Pokemon.PokemonName)
	}

	return nil
}