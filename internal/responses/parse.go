package responses

import (
	"encoding/json"
)

func Parse(res []byte) (interface{}, error) {
	var response interface{}

	err := json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
