package commands


func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
		},
		"help": {
			Name:        "help",
			Description: "Displays the help",
		},
		"map": {
			Name:        "map",
			Description: "displays all next 20 locations",
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays all previous 20 locations",
		},
		"explore": {
			Name: "explore",
			Description: "displays list of pokemon of the region given in argument",
			
		},
	}
}
