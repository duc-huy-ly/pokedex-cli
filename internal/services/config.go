package services

import (
	"time"
)

const (
	PokemonAPIEndpoint     = "https://pokeapi.co/api/v2/"
	DefaultTimeoutDuration = 10 * time.Second // seconds
)

var CurrentState = ProgramCurrentState{
	Cache:       *NewCache(DefaultTimeoutDuration),
	CurrentPage: PokemonAPIEndpoint,
}

type ProgramCurrentState struct {
	CurrentPage  string
	PreviousPage string
	Cache        Pokecache
}
