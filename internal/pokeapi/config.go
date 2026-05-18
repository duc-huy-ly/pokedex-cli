package pokeapi

import (
	"pokedex_cli/internal/pokecache"
	"time"
)

const (
	pokemonAPIEndpoint    = "https://pokeapi.co/api/v2/"
	timeoutDuration = 10 // seconds
)

var LocalConfig = Config{
	Cache: *pokecache.NewCache(timeoutDuration * time.Second),
}

type Config struct {
	NextPageUrl     string
	PreviousPageUrl string
	Cache           pokecache.Cache
}
