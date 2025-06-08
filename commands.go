package main

type cliCommand struct {
	name 	  		string
	description 	string
	callback 		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon_name or pokemon_ID>",
			description: "Catch a Pokemon by ID or name.",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_name or location_ID>",
			description: "Explore a location area in the Pokemon world by ID or name.",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback: commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Displays details about the pokemon if the player has caught the pokemon before.",
			callback: commandInspect,
		},
		"map": {
			name:        "map",
			description: "Lists the names of 20 location areas in the Pokemon world. Subsequent calls will show the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the names of the previous 20 location areas in the Pokemon world, if available.",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemon the player has caught",
			callback:    commandPokedex,
		},
	}
}