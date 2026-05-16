package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var mapState = config{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

type LocationAreaEndpoint struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	if c.Next == "" {
		return fmt.Errorf("No link to make the GET request.\n")
	}
	resp, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("error encountered at commandMap() : %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %d and \n%s\n", resp.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("Error reading from the response body : %v\n", err)
	}
	request := LocationAreaEndpoint{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return fmt.Errorf("Error unmarshaling response body, got : %v\n", err)
	}
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	if request.Next != "" {
		// config.Previous <- config.Next
		// config.Next <- request.Next
		setUrl(c, c.Next, request.Next)
	}
	return nil
}

func setUrl(c *config, previous string, next string) {
	c.Previous = previous
	c.Next = next
}

func commandMapB(c *config) error {
	if c.Previous == "" {
		return fmt.Errorf("No link to make the GET request.\n")
	}
	resp, err := http.Get(c.Previous)
	if err != nil {
		return fmt.Errorf("error encountered at commandMap() : %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %d and \n%s\n", resp.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("Error reading from the response body : %v\n", err)
	}
	request := LocationAreaEndpoint{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return fmt.Errorf("Error unmarshaling response body, got : %v\n", err)
	}
	for _, result := range request.Results {
		fmt.Println(result.Name)
	}
	if request.Previous != "" {
		setUrl(c, request.Previous, c.Previous)
	}
	return nil
}
