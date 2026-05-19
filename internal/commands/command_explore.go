package commands

import (
	"fmt"
	"pokedex_cli/internal/services"
)

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

