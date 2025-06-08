package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, unusedArgument ...string) error {
	if len(c.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}