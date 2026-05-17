package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(_interval time.Duration) *Cache {
	result := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	go result.reapLoop(_interval)
	return &result
}

type Cache struct {
	Entries map[string]cacheEntry
	mu      sync.Mutex
}

func (_cache *Cache) Add(key string, val []byte) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	_cache.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (_cache *Cache) Get(key string) ([]byte, bool) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	result, exists := _cache.Entries[key]
	if !exists {
		return nil, false
	}
	return result.Val, true
}

func (_cache *Cache) reapLoop(interval time.Duration) {
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
