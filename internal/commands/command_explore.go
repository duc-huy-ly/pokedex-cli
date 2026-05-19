package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

// "fmt"
// services "pokedex_cli/internal/services"

type commandExplore struct {
	CliCommand
	Argument []string
	State    *services.ProgramStateStruct
}

func NewCommandExplore(command []string) *commandExplore {
	return &commandExplore{
		CliCommand: CliCommand{
			Name:        "explore",
			Description: "given a location from the 'map' command, displays the list of available pokemon from that location",
		},
		State: &services.CurrentState,
		Argument: command,
	}
}

func (command *commandExplore) Execute() error {
	if len(command.Argument) == 0 || command.Argument[0] == "" {
		return fmt.Errorf("Error : no location given\n")
	}
	url := services.DefaultLocationUrl + "/" + command.Argument[0]	
	locationInfo, err := command.State.Cache.LocationInformation(url)
	if err != nil {
		return err
	}
	for _, pokemon := range locationInfo.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)	
	}
	return nil
}

// func CommandExplore(cfg *services.Config, args []string) error {
// if len(args) == 0 || args[0] == "" {
// 	return fmt.Errorf("No region given to explore")
// }
// locationUrl := services.DefaultLocationUrl + "/" + args[0]
// // with the cache
// cacheData, exists := cfg.Cache.Get(locationUrl)
// if exists {
// 	decodedResponse, err := services.UnmarshalToLocationInfo(cacheData)
// 	if err != nil {
// 		return fmt.Errorf("Error decoding the location info from the response. %v\n", err)
// 	}
// 	for _, pokemon := range decodedResponse.PokemonEncounters {
// 		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
// 	}
// 	fmt.Printf("-------------Taken from cache ------------\n")
// 	return nil

// }
// // make request to pokemon api endpoint else not in cahce
// client := services.NewApiCalls(services.DefaultTimeoutDuration)
// resp, err := client.SendRequest("GET", locationUrl)
// if err != nil {
// 	return fmt.Errorf("Command explore : sending request to api failed. %v\n", err)
// }

// decodedResponse, err := services.UnmarshalToLocationInfo(resp)
// if err != nil {
// 	return fmt.Errorf("Error decoding the location info from the response. %v\n", err)
// }

// for _, pokemon := range decodedResponse.PokemonEncounters {
// 	fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
// }

// cfg.Cache.Add(locationUrl, resp)
// return nil
// }
