package services

type PokeapiService interface {
	LocationAreas(string) (ListOfLocations, error)
	AreaInformation(string) (LocationInfo, error)
}
