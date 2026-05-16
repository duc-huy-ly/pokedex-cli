package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

)
type LocationAreaEndpoint struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
func commandMap() error {
	resp, err :=	http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("error encountered at commandMap() : %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and \n%s\n", resp.StatusCode, body)	
	}
	if err != nil {
		log.Fatal(err)
	}
	location := LocationAreaEndpoint{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return fmt.Errorf("Error unmarshaling response body, got : %v\n", err)
	}
	for _, result := range location.Results {
		fmt.Println(result.Name)	
	}

	return nil
}