package requests

import (
	"fmt"
	"io"
	"net/http"
)

func GetWithQuery(endpoint string, query string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, API_URL+endpoint, nil)
	req.URL.RawQuery = query

	if err != nil {
		return nil, err
	}

	return executeRequest(req)

}

func Get(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, API_URL+endpoint, nil)

	if err != nil {
		return nil, err
	}

	return executeRequest(req)
}

func executeRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-APIKEY", "<PLACEHOLDER>")
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
