package commands

import (
	"fmt"

	"pokedex_cli/internal/pokeapi"
)

func CommandMap(cfg *pokeapi.Config) error {
	url := cfg.NextPageUrl
	if url == "" {
		return fmt.Errorf("No link to make the GET request.\n")
	}
	// Cache first

	data, exists := cfg.Cache.Get(url)
	if exists {
		locations, err := pokeapi.Convert(data)
		if err != nil {
			return err
		}
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		cfg.PreviousPageUrl = cfg.NextPageUrl
		cfg.NextPageUrl = locations.Next
		fmt.Println("#########################")
		fmt.Println("Data recovered from cache")
		fmt.Println("#########################")
		return nil
	}

	// case not in case, do the api call
	request, err := pokeapi.ListLocations(cfg.NextPageUrl)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	if request.Next != "" {
		cfg.PreviousPageUrl = cfg.NextPageUrl
		cfg.NextPageUrl = request.Next
	}
	// update the cache
	return nil
}
