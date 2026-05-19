package main

import "pokedex_cli/internal/services"

func main() {
	StartRepl(&services.LocalConfig)
}
