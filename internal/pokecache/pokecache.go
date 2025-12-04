package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	results map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
	ticker *time.Ticker
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {

	c &cache {
		results: make(map[string]cacheEntry)
		interval: interval
	}

	go cache.reapLoop()
	return c
}

func (c Cache) Add(key string, value []byte) {
	c.Results[key].val = value
}

func (c Cache) get(key string) ([]byte, bool) {
	stored, ok:=c.Results[key]
	if !ok {
		return 0, false
	}
	return stored, true
}

func (c *cache) reapLoop() {

	ticker:=time.NewTicker(c.interval)
	defer ticker.Stop()


	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.results {
				age:=time.Since(entry.createdAt)
				if age > c.interval {
					delete(c.results, key)
				}
			}
		c.mu.Unlock()
	}
	for key, entry := range c.results {
		age:=time.Since(entry.createdAt)
		if age > c.interval {
			delete(c.results, key)
		}
	}

}
