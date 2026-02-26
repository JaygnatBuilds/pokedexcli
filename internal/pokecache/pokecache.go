package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]CacheEntry
	mutex *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	cache := &Cache{
		data:  make(map[string]CacheEntry),
		mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {

	cache.mutex.Lock()

	cache.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	defer cache.mutex.Unlock()

}

func (cache *Cache) Get(key string) ([]byte, bool) {

	cache.mutex.Lock()

	value, ok := cache.data[key]

	defer cache.mutex.Unlock()

	return value.val, ok

}

func (cache *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {

		cache.mutex.Lock()
		defer cache.mutex.Unlock()

		for key, val := range cache.data {

			time := time.Now()
			if time.Sub(val.createdAt) > interval {
				delete(cache.data, key)
			}
		}

	}

}
