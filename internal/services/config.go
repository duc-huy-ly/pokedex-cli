package services

import (
	"time"
)

const (
	PokemonAPIEndpoint     = "https://pokeapi.co/api/v2/"
	DefaultTimeoutDuration = 10 * time.Second // seconds
)

var LocalConfig = Config{
	Cache: *NewCache(DefaultTimeoutDuration),
}

type Config struct {
	NextPageUrl     string
	PreviousPageUrl string
	Cache           Pokecache
}
