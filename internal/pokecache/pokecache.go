package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Results map[string]cacheEntry {
		createdAt time.Time
		val []byte
	}
	mu sync.Mutex
}

func (c Cache) Add(key string, value []byte) {
 c.Results[key].val = value
}
