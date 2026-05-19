package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

type CommandMap struct {
	CliCommand
	state *services.ProgramStateStruct
}

func NewCommandMap() *CommandMap {
	return &CommandMap{
		CliCommand: CliCommand{
			Name:        "map",
			Description: "Displays the 20 locations of the next url page",
		},
		state: &services.CurrentState,
	}
}

func (c *CommandMap) Execute() error {
	// Delegates work to the cache
	url := c.state.NextPage
	if url == services.PokemonAPIEndpoint || url == "" {
		url = services.DefaultLocationUrl
	}
	locations, err := c.state.Cache.LocationAreas(url)
	if err != nil {
		return fmt.Errorf("Something went wrong : %v\n", err)
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	c.state.NextPage = locations.Next
	c.state.PreviousPage = locations.Previous
	return nil
}
