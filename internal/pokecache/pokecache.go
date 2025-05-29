package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries 	map[string]cacheEntry
	mu 			sync.RWMutex
	interval 	time.Duration
}

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: make(map[string]cacheEntry),
		mu: 	 sync.RWMutex{},
		interval: interval,
	}

	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	if time.Since(entry.createdAt) > c.interval {
		delete(c.entries, key)
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <- ticker.C:
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}
}