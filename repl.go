package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex_cli/internal/commands"
	"pokedex_cli/internal/services"
	"strings"
)

func StartRepl(cfg *services.ProgramStateStruct) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nPokedex > ")
		_ = scanner.Scan()
		tokens := CleanInput(scanner.Text())
		if len(tokens) == 0 {
			continue
		}
		command, exists := commands.GetCommands(tokens[1:])[tokens[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.Execute()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
