package services

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Implements the PokeService interface
type Pokecache struct {
	service PokeApiServiceImpl
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(_interval time.Duration) *Pokecache {
	result := Pokecache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	go result.reapLoop(_interval)
	return &result
}

// The method will check if there the requested data exists locallly (cache). If not makes a http GET request to
// the pokemonAPIV2, adds the data to the cache before Unmarshaling and returning the result
func (cache *Pokecache) LocationAreas(url string) (ListOfLocations, error) {
	localData, exists := cache.Get(url)
	if exists {
		fmt.Println("------CACHE-----")
		result, err := UnmarshalDataToListOfLocation(localData)
		if err != nil {
			return ListOfLocations{}, fmt.Errorf("Error encountered while unmarshaling data: %v\n", err)
		}
		return result, nil

	}
	// write to cache before returning the result of the request
	dataFromServer, err := cache.service.MakeRequest("GET", url)
	if err != nil {
		return ListOfLocations{}, err
	}
	cache.Add(url, dataFromServer)
	result, err := UnmarshalDataToListOfLocation(dataFromServer)
	if err != nil {
		return ListOfLocations{}, fmt.Errorf("Error unmarshaling the data to ListOfLocation : %v\n", err)
	}
	return result, nil
}

func (cache *Pokecache) LocationInformation(url string) (LocationInfoStruct, error) {
	localData, exists := cache.Get(url)
	if exists {
		result, err := UnmarshalToLocationInfo(localData)
		if err != nil {
			return LocationInfoStruct{}, fmt.Errorf("Error unmarshaling data to LocationInformation")
		}
		return result, nil
	}
	dataFromServer, err := cache.service.MakeRequest("GET", url)
	if err != nil {
		return LocationInfoStruct{}, err
	}
	// Don't forget to update the cache
	cache.Add(url, dataFromServer)
	result, err := UnmarshalToLocationInfo(dataFromServer)
	if err != nil {
		return LocationInfoStruct{}, err
	}
	return result, nil
}

func (cache *Pokecache) GetPokemon(url string) (PokemonStruct, error) {
	localData, exists := cache.Get(url)
	if exists {
		pokemon := PokemonStruct{}
		err := json.Unmarshal(localData, &pokemon)
		if err != nil {
			return pokemon, fmt.Errorf("Error unmarshaling : %v\n", err)
		}
		return pokemon, nil
	}
	dataFromServer, err := cache.service.MakeRequest("GET", url)
	if err != nil {
		return PokemonStruct{}, fmt.Errorf("Error fetching resource from server. Check spelling of pokemon ? Got: %v\n", err)
	}
	pokemon := PokemonStruct{}
	err = json.Unmarshal(dataFromServer, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("Error unmarshaling data : %v\n", err)
	}
	return pokemon, nil


}

// Updates the map of data safely (taking into account concurrency )
func (_cache *Pokecache) Add(key string, val []byte) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	_cache.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (_cache *Pokecache) Get(key string) ([]byte, bool) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	result, exists := _cache.entries[key]
	if !exists {
		return nil, false
	}
	return result.val, true
}

// Every interval, removes elements from the cache that are too old
func (_cache *Pokecache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		_cache.mu.Lock()
		for key, cacheEntry := range _cache.entries {
			if cacheEntry.createdAt.Add(interval).Before(time.Now()) {
				// expired
				delete(_cache.entries, key)
			}
		}
		_cache.mu.Unlock()
	}
}
