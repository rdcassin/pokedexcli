package main

import (
	"time"
	"github.com/rdcassin/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &config{
		pokeapiClient: pokeClient,
		Next: 	   nil,
		Previous: nil,
	}
	repl(c)
}