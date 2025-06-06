package main

import (
	"fmt"
	"errors"
)

func commandMapb(c *config, unusedArgument string) error {
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
		id, err := getLocationAreaID(location.URL)
		if err != nil {
			return err
		} else {
			fmt.Printf("%v. %v\n", id, location.Name)
		}
	}
	return nil
}