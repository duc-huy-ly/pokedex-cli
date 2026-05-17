package pokeapi

import (
	"net/http"
	"pokedex_cli/internal/pokecache"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

var LocalConfig = Config{
	NextPageUrl:     "https://pokeapi.co/api/v2/location-area/",
	PreviousPageUrl: "",
	Client: http.Client{
		Timeout: 10,
	},
	Cache: *pokecache.NewCache(30 * time.Second),
}

type Config struct {
	NextPageUrl     string
	PreviousPageUrl string
	Client          http.Client
	Cache           pokecache.Cache
}
