package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()
		token := CleanInput(scanner.Text())
		if len(token) == 0 {
			continue
		}
		command, exists := GetCommands()[token[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays the help",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "displays all next 20 locations",
			callback: func() error {
				return CommandMap(&MapState)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "displays all previous 20 locations",
			callback: func() error {
				return CommandMapB(&MapState)
			},
		},
	}
}
