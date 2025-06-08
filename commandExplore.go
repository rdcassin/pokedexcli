package main

import (
	"fmt"
	"errors"
)

func commandExplore(c *config, areaIDOrName ...string) error {
	if len(areaIDOrName) == 0 {
		return errors.New("please provide an area ID or name to explore")
	}

	IDOrName := areaIDOrName[0]

	areaInfo, err := c.pokeapiClient.GetAreaInfo(IDOrName)
	if err != nil {
		return err
	}

	possiblePokemon := areaInfo.PokemonEncounters

	if len(possiblePokemon) == 0 {
		fmt.Printf("No Pokemon found in area: %s\n", IDOrName)
		return nil
	}

	fmt.Printf("Exploring %s...\n", IDOrName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range possiblePokemon {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}