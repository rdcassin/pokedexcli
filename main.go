package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		text := cleanInput(input)
		if len(text) > 0 {
		fmt.Printf("Your command was: %s\n", text[0])
		} else {
			fmt.Println("No command entered.")
		}
	}
}

func cleanInput(text string) []string {
	lowerCaseText:= strings.ToLower(text)
	splitText := strings.Fields(lowerCaseText)
	return splitText
}