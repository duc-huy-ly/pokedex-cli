package commands

import (
	"fmt"
	"time"

	"pokedex_cli/internal/pokeapi"
)

func CommandMap(cfg *pokeapi.Config, args []string) error {
	url := cfg.NextPageUrl
	if url == "" {
		url = pokeapi.DefaultLocationUrl
	}
	// Cache first
	data, exists := cfg.Cache.Get(url)
	if exists {
		locations, err := pokeapi.UnmarshalLocation(data)
		if err != nil {
			return err
		}
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		cfg.PreviousPageUrl = locations.Previous
		cfg.NextPageUrl = locations.Next
		fmt.Println("#########################")
		fmt.Println("Data recovered from cache")
		fmt.Println("#########################")
		return nil
	}

	// case not in case, do the api call
	client := pokeapi.NewClient(5 * time.Second)
	request, err := pokeapi.MakeRequest(*client, url)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	locations, err := pokeapi.UnmarshalLocation(request)
	if err != nil {
		return err
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	cfg.NextPageUrl = locations.Next
	cfg.PreviousPageUrl = locations.Previous

	cfg.Cache.Add(url, request)
	return nil
}
