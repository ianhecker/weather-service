package client

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpclient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpclient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (client *Client) NewRequest(method string, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func (client *Client) Do(request *http.Request) ([]byte, error) {
	response, err := client.httpclient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
