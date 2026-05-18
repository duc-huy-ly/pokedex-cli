package commands

import (
	"fmt"
	"os"
	"pokedex_cli/internal/pokeapi"
)

func CommandExit(cfg *pokeapi.Config, args[]string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("os.Exit didn't close properly")
}