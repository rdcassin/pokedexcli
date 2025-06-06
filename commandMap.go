package main

import (
	"fmt"
)

func commandMap(c *config, unusedArgument string) error {
	locations, err := c.pokeapiClient.GetMap(c.Next)
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
