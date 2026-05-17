package pokeapi

import (
	"pokedex_cli/internal/pokecache"
	"time"
)

var MapState = Config{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

var MyCache = pokecache.NewCache(30 * time.Second)

type Config struct {
	Next     string
	Previous string
}
