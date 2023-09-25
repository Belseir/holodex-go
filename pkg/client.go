package pkg

import (
	"sync"
)

type HolodexClient struct {
	apiKey string
}

var client HolodexClient
var once sync.Once

func NewHolodexClient(apiKey string) HolodexClient {
	once.Do(func() {
		client = HolodexClient{
			apiKey: apiKey,
		}
	})

	return client
}
