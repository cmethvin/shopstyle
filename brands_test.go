package shopstyle_test

import (
	"github.com/cmethvin/shopstyle"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestClient_GetBrands(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		status int
		error  bool
		brands []shopstyle.Brand
	}{
		{
			name:   "test successful request",
			json:   `{"brands":[{"id":"46074","name":"BARBOUR X ALEXA CHUNG"},{"id":"37901","name":"Convenience Concepts"}]}`,
			status: 200,
			error:  false,
			brands: []shopstyle.Brand{
				{ID: "46074", Name: "BARBOUR X ALEXA CHUNG"},
				{ID: "37901", Name: "Convenience Concepts"},
			},
		},
		{
			name:   "test failed request",
			json:   `{"errorCode":401,"errorName":"MissingPartnerId","errorMessage":"pid is required","detailedErrors":[],"errorId":"UXQR0L","headers":{}}`,
			status: 401,
			error:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHttpClient := &MockHttpClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					var response http.Response

					url := req.URL
					assert.Equal(t, "https", url.Scheme)
					assert.Equal(t, "api.shopstyle.com", url.Host)
					assert.Equal(t, "/api/v2/brands", url.Path)

					query := url.Query()
					assert.Equal(t, "fakekey", query.Get("pid"))

					response.Body = io.NopCloser(strings.NewReader(tt.json))
					response.StatusCode = tt.status

					return &response, nil
				},
			}

			client := shopstyle.New("fakekey", shopstyle.WithHttpClient(mockHttpClient))
			brands, err := client.GetBrands()

			assert.True(t, mockHttpClient.DoCalled, "expected http request")
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.brands != nil {
				assert.Equal(t, tt.brands, brands.Brands)
			}
		})
	}
}
