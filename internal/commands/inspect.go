package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

type inspectCommand struct {
	CliCommand
	args  []string
	state *services.ProgramStateStruct
}

func NewInspectCommand(_args []string) *inspectCommand {
	return &inspectCommand{
		CliCommand: CliCommand{
			Name:        "inspect",
			Description: "displays information about the pokemon if caught",
		},
		args:  _args,
		state: &services.CurrentState,
	}
}

func (this *inspectCommand) Execute() error {
	if len(this.args) == 0 || this.args[0] == "" {
		return fmt.Errorf("No argument given")
	}
	pokemonName := this.args[0]
	pokemon, exists := this.state.Pokedex[pokemonName]
	if !exists {
		fmt.Printf("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil

}
