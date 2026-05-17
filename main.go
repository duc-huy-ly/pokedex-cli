package main

import "pokedex_cli/internal/pokeapi"

func main() {
	StartRepl(&pokeapi.LocalConfig)
}
