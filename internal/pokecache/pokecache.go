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

	c := &Cache {
		results: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.results[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	stored, ok:=c.results[key]
	if !ok {
		return nil, false
	}
	return stored.val, true
}

func (c *Cache) reapLoop() {

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
