package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	store    map[string]CacheEntry
	mu       sync.RWMutex
	stopChan chan struct{}
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		interval: interval,
		store:    make(map[string]CacheEntry),
		stopChan: make(chan struct{}),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.store[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func c(c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.reap()
		case <-c.stopChan:
			return 
		}
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	cutoff := time.Now().Add(-c.interval) 
	for key, entry := range c.store {
		if entry.createdAt.Before(cutoff) {
			delete(c.store, key)
		}
	}
}

func (c *Cache) Shutdown() {
	close(c.stopChan)
}