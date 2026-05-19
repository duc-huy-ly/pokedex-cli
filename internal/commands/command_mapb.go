package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

// "fmt"
// "pokedex_cli/internal/pokeapi"
// "time"

type CommandMapB struct {
	CliCommand
	*services.ProgramCurrentState
}

func NewCommandMapB() *CommandMapB {
	return &CommandMapB{
		CliCommand: CliCommand{
			Name:        "mapb",
			Description: "Displays the 20 locations of the previous area page",
		},
		ProgramCurrentState: &services.CurrentState,
	}
}

func (c CommandMapB) Execute() error {
	url := c.ProgramCurrentState.PreviousPage
	if url == services.PokemonAPIEndpoint || url == "" {
		url = services.DefaultLocationUrl
	}
	locations, err := c.ProgramCurrentState.Cache.LocationAreas(url)
	if err != nil {
		return fmt.Errorf("Something went wrong : %v\n", err)
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	c.CurrentPage = locations.Next
	c.PreviousPage = locations.Previous
	return nil
}
