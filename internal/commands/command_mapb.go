package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

type CommandMapB struct {
	CliCommand
	State *services.ProgramStateStruct
}

func NewCommandMapB() *CommandMapB {
	return &CommandMapB{
		CliCommand: CliCommand{
			Name:        "mapb",
			Description: "Displays the 20 locations of the previous area page",
		},
		State: &services.CurrentState,
	}
}

func (c CommandMapB) Execute() error {
	url := c.State.PreviousPage
	if url == services.PokemonAPIEndpoint || url == "" {
		url = services.DefaultLocationUrl
	}
	locations, err := c.State.Cache.LocationAreas(url)
	if err != nil {
		return fmt.Errorf("Something went wrong : %v\n", err)
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	c.State.NextPage = locations.Next
	c.State.PreviousPage = locations.Previous
	return nil
}
