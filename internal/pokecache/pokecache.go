package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	results map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
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

	//start reapLoop here
	//use a ticker to determine how many seconds have passed between each call
	//eg: ticker is reset at the top of every call to know how long it has been since last call

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

func (c cache) reapLoop() {

}
