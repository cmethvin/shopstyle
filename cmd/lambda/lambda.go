package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cmethvin/shopstyle"
	"log"
)

const apiKey = "shopstyle"

func Handler() error {
	client := shopstyle.New(apiKey)

	log.Println("retrieving brands from shopstyle")

	brands, err := client.GetBrands()
	if err != nil {
		return fmt.Errorf("error retrieving brands: %w", err)
	}

	log.Printf("retrieved %d brands", len(brands.Brands))

	return nil
}

func main() {
	lambda.Start(Handler)
}
