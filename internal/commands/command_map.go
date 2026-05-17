package commands

import (
	"fmt"

	"pokedex_cli/internal/pokeapi"
)

func CommandMap(c *pokeapi.Config) error {
	url := c.Next
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
		c.Previous = c.Next
		c.Next = locations.Next
		fmt.Println("#########################")
		fmt.Println("Data recovered from cache")
		fmt.Println("#########################")
		return nil
	}

	// case not in case, do the api call
	request, err := pokeapi.ListLocations(c.Next)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	if request.Next != "" {
		c.Previous = c.Next
		c.Next = request.Next
	}
	// update the cache 
	return nil
}
