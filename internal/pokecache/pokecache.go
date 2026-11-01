package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
type CacheMap struct {
	mu *sync.Mutex
	caches map[string]cacheEntry
}
// var datacache CacheMap

func NewCache(interval time.Duration) CacheMap {
	datacache := CacheMap{
		mu: &sync.Mutex{},
		caches: make(map[string]cacheEntry),
	}
	//defer ticker.Stop()
	go datacache.reapLoop(interval)
	// go func() {
	// 	for t := range datacache.this_ticker.C {
	// 		datacache.ReapLoop(t)
	// 	}
	// }()
	return datacache
}

// func (datacache *CacheMap) EndCache() {
// 	cache := *datacache
// 	cache.this_ticker.Stop()
// 	// datacache = nil
// 	*datacache = cache
// }

func (datacache *CacheMap) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		datacache.reap(time.Now().UTC(), interval)
	}
}

func (datacache *CacheMap) reap(now time.Time, last time.Duration) {
	datacache.mu.Lock()
	defer datacache.mu.Unlock()
	for k, v := range datacache.caches {
		if v.createdAt.Before(now.Add(-last)) {
			delete(datacache.caches, k)
		}
	}
}

func (datacache *CacheMap) Add(name string, data []byte) {
	datacache.mu.Lock()
	defer datacache.mu.Unlock()
	datacache.caches[name] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: data,
	}
}

func (datacache *CacheMap) Get(name string) ([]byte, bool) {
	datacache.mu.Lock()
	defer datacache.mu.Unlock()
	entry, ok := datacache.caches[name]
	return entry.val, ok
}
