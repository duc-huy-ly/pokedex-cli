package commands

import (
	"fmt"
	"pokedex_cli/internal/pokeapi"
)

func CommandMapB(c *pokeapi.Config) error {
	url := c.Previous
	if url == "" {
		return fmt.Errorf("No link to make the GET request.\n")
	}
	request, err := pokeapi.ListLocations(url)

	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	if request.Previous != "" {
		c.Previous = request.Previous
		c.Next = c.Previous
	}
	return nil
}
