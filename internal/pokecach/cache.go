package pokecach

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}


type cacheEntry struct {
	rawresponse []byte
	created time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.LoopReep(interval)
	return c
}

func (c *Cache) Add(key string, rwRes []byte){
	c.cache[key] = cacheEntry{
		rawresponse: rwRes,
		created: time.Now().UTC(),
	}
}

func (c *Cache) LoopReep (interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.AddEr(interval)
	}
}

func (c *Cache) AddEr(interval time.Duration){
	intervaleft := time.Now().UTC().Add(-interval)
	for i, n := range c.cache {
		if n.created.Before(intervaleft){
			delete(c.cache, i)
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	v, o := c.cache[key]
	return v.rawresponse, o
}