package shopstyle

import (
	"net/http"
	"time"
)

const apiUrl = "https://www.shopstyle.com/api/v2/"

type ClientOption func(client *Client)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HttpClient
	key        string
}

func New(key string, opts ...ClientOption) *Client {
	defaultHttpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	client := &Client{
		httpClient: defaultHttpClient,
		key:        key,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithHttpClient(httpClient HttpClient) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func (c *Client) buildRequest(resource string, params map[string][]string) *http.Request {
	req, _ := http.NewRequest("GET", apiUrl+resource, nil)
	query := req.URL.Query()
	query.Set("pid", c.key)

	for k, values := range params {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()
	return req
}
