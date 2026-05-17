package commands

import "pokedex_cli/internal/pokeapi"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config) error
}
