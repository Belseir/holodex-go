package pkg

import (
	"fmt"

	"github.com/belseir/holodex-go/internal/requests"
	"github.com/belseir/holodex-go/internal/responses"
	"github.com/belseir/holodex-go/internal/utils"
	"github.com/belseir/holodex-go/pkg/models"
)

type GetVideoOptions struct {
	Query *models.VideoQueryParams
}

func (c *HolodexClient) GetVideos(opts ...GetVideoOptions) ([]models.VideoWithChannel, error) {
	var query string

	if len(opts) > 0 && opts[0].Query != nil {
		query = "/videos" + opts[0].Query.GetQueryString()
	} else {
		query = "/videos"
	}

	cacheKey := utils.Hash(query)

	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached.([]models.VideoWithChannel), nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	videos, err := responses.Parse[[]models.VideoWithChannel](res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = res

	if err != nil {
		return nil, err
	}

	return *videos, nil
}

func (c *HolodexClient) GetVideoMetadata(videoId string) (*models.VideoFull, error) {
	query := fmt.Sprintf("/videos/%s?c=1", videoId)
	cacheKey := utils.Hash(query)

	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached.(*models.VideoFull), nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	videos, err := responses.Parse[models.VideoFull](res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = res

	if err != nil {
		return nil, err
	}

	return videos, nil
}

type GetUpcomingVideosOptions struct {
	Query *models.VideoQueryParams
}

func (c *HolodexClient) GetUpcomingVideos(opts ...GetUpcomingVideosOptions) ([]models.VideoWithChannel, error) {
	var query string

	if len(opts) > 0 && opts[0].Query != nil {
		query = "/live" + opts[0].Query.GetQueryString()
	} else {
		query = "/live"
	}

	cacheKey := utils.Hash(query)

	if cached := c.searchInCache(cacheKey); cached != nil {
		return cached.([]models.VideoWithChannel), nil
	}

	res, err := requests.Get(query, c.apiKey)
	if err != nil {
		return nil, err
	}

	videos, err := responses.Parse[[]models.VideoWithChannel](res)
	if err != nil {
		return nil, err
	}

	c.cache[cacheKey] = res

	if err != nil {
		return nil, err
	}

	return *videos, nil
}
