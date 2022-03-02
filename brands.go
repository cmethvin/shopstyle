package shopstyle

import (
	"encoding/json"
	"fmt"
)

type Brand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BrandsResponse struct {
	Brands []Brand `json:"brands"`
}

func (c *Client) GetBrands() (BrandsResponse, error) {
	var brandsResponse BrandsResponse

	req := c.buildRequest("brands", nil)
	resp, _ := c.httpClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return brandsResponse, fmt.Errorf("received http status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&brandsResponse); err != nil {
		return brandsResponse, err
	}

	return brandsResponse, nil
}
