package commands


func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays the help",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "displays all next 20 locations",
			Callback: 	CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays all previous 20 locations",
			Callback: CommandMapB,
		},
	}
}
