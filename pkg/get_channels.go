package pkg

import (
	"fmt"

	"github.com/belseir/holodex-go/internal/requests"
	"github.com/belseir/holodex-go/internal/responses"
	"github.com/belseir/holodex-go/internal/utils"
	"github.com/belseir/holodex-go/pkg/models"
)

func (c *HolodexClient) GetChannel(channelId string) (*models.Channel, error) {
	query := fmt.Sprintf("/channels/%s", channelId)

	cacheKey := utils.Hash(query)

	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached.(*models.Channel), nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	channel, err := responses.Parse[models.Channel](res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = res

	if err != nil {
		return nil, err
	}

	return channel, nil
}

type GetChannelsOptions struct {
	Query *models.ChannelQueryParams
}

func (c *HolodexClient) GetChannels(opts ...GetChannelsOptions) ([]models.Channel, error) {
	var query string

	if len(opts) > 0 && opts[0].Query != nil {
		query = "/channels" + opts[0].Query.GetQueryString()
	} else {
		query = "/channels"
	}

	cacheKey := utils.Hash(query)

	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached.([]models.Channel), nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	channels, err := responses.Parse[[]models.Channel](res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = res

	if err != nil {
		return nil, err
	}

	return *channels, nil
}
