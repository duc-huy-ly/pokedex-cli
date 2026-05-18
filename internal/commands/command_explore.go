package commands

import (
	"fmt"
	"pokedex_cli/internal/pokeapi"
)


func CommandExplore(cfg *pokeapi.Config, args[]string) error{
	if len(args) == 0 {
		return fmt.Errorf("No region given to explore")
	}
	return nil
}