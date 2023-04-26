package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	timeCheckDelete := time.Now().UTC().Add(-interval)

	for key, cacheEntry := range c.cache {
		if cacheEntry.createdAt.Before(timeCheckDelete) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}
