package pokeapi

type PokeapiService interface {
	LocationAreas(string) []byte
	AreaInformation(string) []byte
}
