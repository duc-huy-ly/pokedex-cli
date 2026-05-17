package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsRequest struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocations(url string) (LocationsRequest, error) {
	request := LocationsRequest{}
	resp, err := http.Get(url)
	if err != nil {
		return request, fmt.Errorf("error encountered at commandMap() : %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return request, fmt.Errorf("Response failed with status code %d and \n%s\n", resp.StatusCode, body)
	}
	if err != nil {
		return request, fmt.Errorf("Error reading from the response body : %v\n", err)
	}
	// update the cache 
	LocalConfig.Cache.Add(url, body)
	return Convert(body)
}

func Convert(data []byte) (LocationsRequest, error) {
	request := LocationsRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		return request, fmt.Errorf("Error unmarshaling response body, got : %v\n", err)
	}
	return request, nil
}
