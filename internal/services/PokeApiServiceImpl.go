package services

import (
	"io"
	"net/http"
	"time"
)

type PokeApiServiceImpl struct {
	client http.Client
}

func NewApiCalls(timeoutDuration time.Duration) *PokeApiServiceImpl {
	return &PokeApiServiceImpl{
		client: http.Client{
			Timeout: timeoutDuration,
		},
	}
}

// Calls the PokeAPIV2 using a http.Client
func (myClient PokeApiServiceImpl) MakeRequest(requestType string, url string) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := myClient.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return body, err
}
