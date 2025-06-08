package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"github.com/rdcassin/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next 		  *string
	Previous 	  *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func repl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		text := cleanInput(input)

		command := text[0]
		commandParameter := []string{}
		if len(text) > 1 {
			commandParameter = text[1:]
		}

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Print("Unknown command\n")
		} else {
			err := cmd.callback(c, commandParameter...)
			if err != nil {
				fmt.Println(err)				
			}
		}
	}
}

func cleanInput(text string) []string {
	lowerCaseText:= strings.ToLower(text)
	splitText := strings.Fields(lowerCaseText)
	return splitText
}