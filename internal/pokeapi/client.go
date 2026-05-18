package pokeapi

import (
	"io"
	"net/http"
	"time"
)

type MyClient struct {
	client http.Client
}

func NewClient(timeoutDuration time.Duration) *MyClient {
	return &MyClient{
		client: http.Client{
			Timeout: timeoutDuration,
		},
	}
}

func (myClient *MyClient) SendRequest(requestType string, url string) ([]byte , error){
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
	return body, nil
}
func MakeRequest(client MyClient, url string) ([]byte, error) {
	data, err := client.SendRequest("GET", url)
	if err != nil {
		return nil, err
	}
	return data, err
}