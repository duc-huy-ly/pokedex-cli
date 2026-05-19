package services

type PokeapiService interface {
	LocationAreas(string) (ListOfLocations, error)
	LocationInformation(string) (LocationInfoStruct, error)
}
