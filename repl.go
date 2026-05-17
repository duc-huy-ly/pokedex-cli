package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"pokedex_cli/internal/commands"
)



func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()
		token := CleanInput(scanner.Text())
		if len(token) == 0 {
			continue
		}
		command, exists := commands.GetCommands()[token[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.Callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

