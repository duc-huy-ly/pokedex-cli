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
	// Cache first
	
	data, exists := pokeapi.MyCache.Get(url)
	if exists {
		locations, err :=pokeapi.Convert(data)
		if err != nil {
			return err
		}
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		c.Previous = locations.Previous
		c.Next = c.Previous
		fmt.Println("#########################")
		fmt.Println("Data recovered from cache")
		fmt.Println("#########################")
		return nil
	}

	// case not in case, do the api call

	request, err := pokeapi.ListLocations(url)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	// displays the 20 locations 
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	/// update the list of the configuration file
	if request.Previous != "" {
		c.Previous = request.Previous
		c.Next = c.Previous
	}
	return nil
}
