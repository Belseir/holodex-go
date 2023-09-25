package responses

import (
	"encoding/json"

	"github.com/belseir/holodex-go/pkg/models"
)

func ParseResponse[T models.Channel | models.ChannelWithGroup | models.Video | models.VideoFull | models.VideoWithChannel |
	[]models.Channel | []models.ChannelWithGroup | []models.Video | []models.VideoFull | []models.VideoWithChannel](res []byte) (*T, error) {
	var response *T

	err := json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
