package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

type CommandMap struct {
	CliCommand
	*services.ProgramCurrentState
}

func NewCommandMap() *CommandMap {
	return &CommandMap{
		CliCommand: CliCommand{
			Name:        "map",
			Description: "Displays the 20 locations of the next url page",
		},
		ProgramCurrentState: &services.CurrentState,
	}
}

func (c *CommandMap) Execute() error {
	// Delegates work to the cache
	url := c.ProgramCurrentState.CurrentPage
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
