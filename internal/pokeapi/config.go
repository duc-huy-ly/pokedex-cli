package pokeapi

import (
	"pokedex_cli/internal/pokecache"
	"time"
)

const (
	PokemonAPIEndpoint    = "https://pokeapi.co/api/v2/"
	DefaultTimeoutDuration = 10 * time.Second// seconds
)

var LocalConfig = Config{
	Cache: *pokecache.NewCache(DefaultTimeoutDuration),
}

type Config struct {
	NextPageUrl     string
	PreviousPageUrl string
	Cache           pokecache.Cache
}
