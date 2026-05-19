package services

import (
	"time"
)

const (
	PokemonAPIEndpoint     = "https://pokeapi.co/api/v2/"
	DefaultTimeoutDuration = 10 * time.Second
)

var CurrentState = ProgramStateStruct{
	Cache:    *NewCache(DefaultTimeoutDuration),
	NextPage: PokemonAPIEndpoint,
	Pokedex: make(map[string]PokemonStruct),
}

type ProgramStateStruct struct {
	NextPage     string
	PreviousPage string
	Cache        Pokecache
	Pokedex map[string]PokemonStruct
}
