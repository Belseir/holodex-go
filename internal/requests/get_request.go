package requests

import (
	"fmt"
	"io"
	"net/http"
)

func Get(query string, apiKey string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, API_URL+query, nil)

	if err != nil {
		return nil, err
	}

	return executeRequest(req, apiKey)
}

func executeRequest(req *http.Request, apiKey string) ([]byte, error) {
	req.Header.Set("X-APIKEY", apiKey)
	raw, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if raw.StatusCode == 403 {
		return nil, fmt.Errorf("Invalid API key provided.")
	}

	response, err := io.ReadAll(raw.Body)

	if err != nil {
		return nil, err
	}

	return response, nil
}
