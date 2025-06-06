package main

import (
	"time"
	"github.com/rdcassin/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	c := &config{
		pokeapiClient: pokeClient,
		Next: 	   nil,
		Previous: nil,
	}
	repl(c)
}