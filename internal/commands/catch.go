package commands

import (
	"fmt"
	"math/rand"
	"pokedex_cli/internal/services"
)

type commandCatch struct {
	CliCommand
	args  []string
	state *services.ProgramStateStruct
}

func NewCommandCatch(_args []string) *commandCatch {
	return &commandCatch{
		CliCommand: CliCommand{
			Name:        "catch",
			Description: "Try to catch a given pokemon",
		},
		args:  _args,
		state: &services.CurrentState,
	}
}

func (command *commandCatch) Execute() error {
	if len(command.args) == 0 || command.args[0] == "" {
		return fmt.Errorf("No pokemon name given\n")
	}
	name := command.args[0]
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + name
	pokemon, err := command.state.Cache.GetPokemon(endpoint)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

	// calculate how hard it would be to catch the pokemon
	// fmt.Println(pokemon.BaseExperience)
	caught := calculatePokemonCatchChance(pokemon)
	if !caught {
		fmt.Printf("%v broke out!\n", name)
	} else {
		fmt.Printf("%v was caught!\n", name)
		command.state.Pokedex[name] = pokemon
	}
	return nil

}

// this function could be more intricate, for now it will do
func calculatePokemonCatchChance(pokemon services.PokemonStruct) bool {
	var chance float64 = 60.0 / (100.0 + float64(pokemon.BaseExperience))
	value := rand.Float64()
	fmt.Printf("catch chance : %v\nrand value: %v\n",chance, value)
	return value <= chance
}
