package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, pokemonName ...string) error {
	if len(pokemonName) == 0 {
		return errors.New("please enter a pokemon name")
	}

	target := pokemonName[0]

	pokemon, exists := c.caughtPokemon[target]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Details.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Details.Name)
	}

	return nil
}