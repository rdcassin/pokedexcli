package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		text := cleanInput(input)
		command := text[0]

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Print("Unknown command\n")
		} else {
			cmd.callback()
		}
	}
}

func cleanInput(text string) []string {
	lowerCaseText:= strings.ToLower(text)
	splitText := strings.Fields(lowerCaseText)
	return splitText
}