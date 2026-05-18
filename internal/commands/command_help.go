package commands

import (
	"fmt"
	"pokedex_cli/internal/pokeapi"
)

func CommandHelp(cfg *pokeapi.Config, args[]string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
