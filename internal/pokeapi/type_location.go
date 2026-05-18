package pokeapi

import (
	"encoding/json"
	"fmt"
)

const DefaultLocationUrl = "https://pokeapi.co/api/v2/location-area/"

type LocationStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}



func UnmarshalLocation(data []byte) (LocationStruct, error) {
	request := LocationStruct{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		return request, fmt.Errorf("Error unmarshaling response body, got : %v\n", err)
	}
	return request, nil
}
