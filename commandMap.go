package main

import (
	"fmt"
)

func commandMap(c *config) error {
	locations, err := c.pokeapiClient.GetMap(c.Next)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

