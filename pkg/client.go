package pkg

import (
	"sync"

	"github.com/belseir/holodex-go/internal/requests"
	"github.com/belseir/holodex-go/internal/responses"
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

func (c *HolodexClient) executeGetRequest(query string, cacheKey string) (interface{}, error) {
	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached, nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	parsed, err := responses.Parse(res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = parsed

	return parsed, nil
}
