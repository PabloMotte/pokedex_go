package pokeapi

import (
	"net/http"
	"time"

	"github.com/pablomotte/pokedex_go/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.CacheMap
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
