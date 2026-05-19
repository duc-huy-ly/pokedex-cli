package services

import (
	"sync"
	"time"
)

// Implements the PokeService interface
type Pokecache struct {
	Service PokeApiServiceImpl
	Entries map[string]cacheEntry
	mu      sync.Mutex
}
type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(_interval time.Duration) *Pokecache {
	result := Pokecache{
		Entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	go result.reapLoop(_interval)
	return &result
}

func (cache *Pokecache) LocationAreas(url string) []byte {
	return nil
}

func (_cache *Pokecache) Add(key string, val []byte) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	_cache.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (_cache *Pokecache) Get(key string) ([]byte, bool) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	result, exists := _cache.Entries[key]
	if !exists {
		return nil, false
	}
	return result.Val, true
}

func (_cache *Pokecache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		_cache.mu.Lock()
		for key, cacheEntry := range _cache.Entries {
			if cacheEntry.CreatedAt.Add(interval).Before(time.Now()) {
				// expired
				delete(_cache.Entries, key)
			}
		}
		_cache.mu.Unlock()
	}
}
