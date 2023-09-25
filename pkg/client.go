package pkg

import (
	"sync"
)

type HolodexClient struct {
	apiKey string
	cache  map[string]interface{}
}

var client HolodexClient
var once sync.Once

func NewHolodexClient(apiKey string) HolodexClient {
	once.Do(func() {
		client = HolodexClient{
			apiKey: apiKey,
			cache:  make(map[string]interface{}),
		}
	})

	return client
}

func (c *HolodexClient) searchInCache(hash string) interface{} {
	cached, inCache := c.cache[hash]

	if inCache {
		return cached
	}

	return nil
}
