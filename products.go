package shopstyle

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const categoryAll = "clothes-shoes-and-jewelry"

type ProductsOption func(params map[string][]string)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"unbrandedName"`
	Brand Brand  `json:"brand"`
}

type Metadata struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ProductsResponse struct {
	Metadata Metadata  `json:"metadata"`
	Products []Product `json:"products"`
}

func (c *Client) GetProducts(limit, offset int, opts ...ProductsOption) (ProductsResponse, error) {
	var response ProductsResponse

	params := map[string][]string{
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
		"cat":    {categoryAll},
	}

	for _, opt := range opts {
		opt(params)
	}

	req := c.buildRequest("products", params)

	resp, _ := c.httpClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return response, fmt.Errorf("received http status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}

func WithBrandFilter(brandID string) ProductsOption {
	return func(params map[string][]string) {
		params["fl"] = append(params["fl"], "b"+brandID)
	}
}
