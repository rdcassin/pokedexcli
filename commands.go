package main

type cliCommand struct {
	name 	  		string
	description 	string
	callback 		func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the names of 20 location areas in the Pokemon world. Subsequent calls will show the next 20 locations.",
			callback:    commandMap,
		}
	}
}