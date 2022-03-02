package main

import (
	"flag"
	"fmt"
	"github.com/cmethvin/shopstyle"
	"log"
	"os"
)

const apiKey = "shopstyle"

var brandID string

func init() {
	flag.StringVar(&brandID, "b", "", "brand id of the products that will be loaded")
	flag.Parse()
}

func main() {
	log.Println(os.Args)

	if brandID == "" {
		log.Fatal("no brand id specified")
	}

	log.Printf("starting batch job with brand %s", brandID)

	client := shopstyle.New(apiKey)

	response, err := client.GetProducts(40, 0, shopstyle.WithBrandFilter(brandID))
	if err != nil {
		log.Fatal(fmt.Errorf("error loading products: %w", err))
	}

	log.Printf("loaded %d products of %d", len(response.Products), response.Metadata.Total)
}
