package main

import (
	"fmt"
	"errors"
)

func commandMapb(c *config) error {
	if c.Previous == nil {
		return errors.New("no previous map available")
	}
	
	locations, err := c.pokeapiClient.GetMap(c.Previous)
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