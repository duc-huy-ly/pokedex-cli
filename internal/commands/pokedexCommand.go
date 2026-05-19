package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

type pokedexCommand struct {
	CliCommand
	state *services.ProgramStateStruct
}

func NewPokedexCommand() *pokedexCommand {
	return &pokedexCommand{
		CliCommand: CliCommand{
			Name: "pokedex",
			Description: "displays all the pokemons caught",
		},
		state: &services.CurrentState,
	}
}

func (this *pokedexCommand) Execute() error {
	fmt.Println("Your Pokedex :")
	for k := range this.state.Pokedex {
		fmt.Printf("- %v\n", k)
	}
	return nil
}
