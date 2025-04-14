package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct { //THIS IS THE CACHE
	cache map[string]cacheEntry //cache is a MAP, strings are URLS, cacheEntry are objects stored within
	mux   sync.Mutex            //automatically initialized on construction, for protecting goroutines
}
type cacheEntry struct { //DATA STRUCTURE FOR THE ENTRIES, HOLDS UTC TIMESTAMP AND THE DATA ENTRY AS SLICE OF BYTES
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache { //initializes new cache, takes the interval at which to clean and returns a POINTER to the new cache, ensuring all methods effect same cache. Also avoids copying the entire cache when returning
	//when other functions modify the cache, the changes will be visible to all code holding the pointer
	c := &Cache{ //& CREATES A POINTER
		cache: make(map[string]cacheEntry), //to this cache being made
	}
	go c.reapLoop(interval) //begin the reaping in its own routine

	return c //return the pointer to the cache, EVEN THOUGH C HAS NO * IT IS A POINTER
}

func (c *Cache) Add(key string, value []byte) {
	//update c with added key val

	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	//if key is in cache, return true else return false
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in reaploop: ", r)
		}
	}()
	timeclock := time.NewTicker(interval)
	defer timeclock.Stop()

	for range timeclock.C {
		c.reap(time.Now().UTC(), interval)

	}
	//find out how to get interval here
	//on every tick, check for caches createdat less than current time - interval
	//delete old entries
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, value := range c.cache {
		if value.createdAt.Before(now.Add(-last)) { //delete entries older than last
			delete(c.cache, key)
		}
	}
}
