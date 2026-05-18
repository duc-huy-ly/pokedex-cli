package commands

import (
	"fmt"
	"pokedex_cli/internal/pokeapi"
)

func CommandExplore(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 || args[0] == "" {
		return fmt.Errorf("No region given to explore")
	}
	locationUrl := pokeapi.DefaultLocationUrl + "/" + args[0]
	// with the cache
	cacheData, exists := cfg.Cache.Get(locationUrl)
	if exists {
		decodedResponse, err := pokeapi.UnmarshalToLocationInfo(cacheData)
		if err != nil {
			return fmt.Errorf("Error decoding the location info from the response. %v\n", err)
		}
		for _, pokemon := range decodedResponse.PokemonEncounters {
			fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
		}
		fmt.Printf("-------------Taken from cache ------------\n")
		return nil

	}
	// make request to pokemon api endpoint else not in cahce
	client := pokeapi.NewClient(pokeapi.DefaultTimeoutDuration)
	resp, err := client.SendRequest("GET", locationUrl)
	if err != nil {
		return fmt.Errorf("Command explore : sending request to api failed. %v\n", err)
	}

	decodedResponse, err := pokeapi.UnmarshalToLocationInfo(resp)
	if err != nil {
		return fmt.Errorf("Error decoding the location info from the response. %v\n", err)
	}

	for _, pokemon := range decodedResponse.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	cfg.Cache.Add(locationUrl, resp)
	return nil
}
