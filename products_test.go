package shopstyle_test

import (
	"bytes"
	"github.com/cmethvin/shopstyle"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestClient_GetProducts(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		status   int
		error    bool
		limit    int
		offset   int
		options  []shopstyle.ProductsOption
		products shopstyle.ProductsResponse
	}{
		{
			name:   "test successful request",
			json:   "../testdata/shopstyle/products.json",
			status: 200,
			error:  false,
			limit:  40,
			offset: 80,
			products: shopstyle.ProductsResponse{
				Metadata: shopstyle.Metadata{
					Limit:  10,
					Offset: 0,
					Total:  6344382,
				},
				Products: []shopstyle.Product{
					{
						ID:   907623839,
						Name: "Ombre Rainbow Pullover",
						Brand: shopstyle.Brand{
							ID:   "444",
							Name: "Oscar de la Renta",
						},
					},
					{
						ID:   898527391,
						Name: "SSENSE Exclusive Pink Mesh Pollen Turtleneck",
						Brand: shopstyle.Brand{
							ID:   "8167",
							Name: "Henrik Vibskov",
						},
					},
				},
			},
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
					assert.Equal(t, "/api/v2/products", url.Path)

					query := url.Query()
					assert.Equal(t, "fakekey", query.Get("pid"))
					assert.Equal(t, strconv.Itoa(tt.limit), query.Get("limit"), "limit doesn't match")
					assert.Equal(t, strconv.Itoa(tt.offset), query.Get("offset"), "offset doesn't match")
					assert.Equal(t, "clothes-shoes-and-jewelry", query.Get("cat"), "category doesn't match")

					b, err := os.ReadFile(tt.json)
					if err != nil {
						log.Println(err)
						t.Fatal("failed to load testdata file")
					}

					response.Body = io.NopCloser(bytes.NewReader(b))
					response.StatusCode = tt.status

					return &response, nil
				},
			}

			client := shopstyle.New("fakekey", shopstyle.WithHttpClient(mockHttpClient))
			products, err := client.GetProducts(tt.limit, tt.offset, tt.options...)

			assert.True(t, mockHttpClient.DoCalled, "expected http request")
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.products, products)
		})
	}
}
