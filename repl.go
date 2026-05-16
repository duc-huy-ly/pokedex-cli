package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()
		token := cleanInput(scanner.Text())
		if len(token) == 0 {
			continue
		}
		command, exists := getCommands()[token[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("os.Exit didn't close properly")
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays the help",
			callback: commandHelp,
		},
	}
}