package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(_interval time.Duration) *cache {
	result := cache{
		Entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	go result.reapLoop(_interval)
	return &result
}

type cache struct {
	Entries  map[string]cacheEntry
	mu       sync.Mutex
}

func (_cache *cache) Add(key string, val []byte) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	_cache.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (_cache *cache) Get(key string) ([]byte, bool) {
	_cache.mu.Lock()
	defer _cache.mu.Unlock()
	result, exists := _cache.Entries[key]
	if !exists {
		return nil, exists
	}
	return result.Val, exists
}

func (_cache *cache) reapLoop(interval time.Duration) {
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
