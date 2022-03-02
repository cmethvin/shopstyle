package main

import (
	"github.com/cmethvin/shopstyle"
	"log"
	"math/rand"
	"os"
	"time"
)

type ProductData struct {
	Start         time.Time
	End           time.Time
	MaxProducts   int
	TotalProducts int
	Errors        int
}

func main() {
	var data ProductData
	data.Start = time.Now()

	log.Println("starting shopstyle data retrieval")

	apiKey := os.Getenv("SHOPSTYLE_KEY")

	client := shopstyle.New(apiKey)

	log.Println("fetching list of brands")

	brands, err := client.GetBrands()
	if err != nil {
		log.Fatal(err)
	}

	totalBrands := len(brands.Brands)

	log.Printf("got %d brands\n", totalBrands)

	for i, brand := range brands.Brands {
		sleep := 5 + rand.Intn(25)
		time.Sleep(time.Duration(sleep) * time.Second)

		log.Printf("fetching products for brand %s, %d of %d\n", brand.Name, i+1, totalBrands)

		products, err := client.GetProducts(0, 0, shopstyle.WithBrandFilter(brand.ID))
		if err != nil {
			data.Errors++
		}

		log.Printf("brand %s has %d products\n", brand.Name, products.Metadata.Total)

		if products.Metadata.Total > data.MaxProducts {
			data.MaxProducts = products.Metadata.Total
		}

		data.TotalProducts += products.Metadata.Total
	}

	log.Printf("data retrieval completed. elapsed time %.0f minutes\n", time.Now().Sub(data.Start).Minutes())
	log.Printf("total products: %d, max products on a brand: %d", data.TotalProducts, data.MaxProducts)
}
