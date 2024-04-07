package model

import (
	pb "github.com/silverflin/go-rpc/proto"
	"slices"
)

func GetProducts() *pb.ProductList {
	productList := &pb.ProductList{Product: make([]*pb.Product, 0)}

	for _, val := range products {
		if !slices.Contains(productList.Product, val.Product) {
			productList.Product = append(productList.Product, val.Product)
		}
	}

	return productList
}

func GetPricesFromProduct(productName string) []*pb.MarketPrice {
	for _, val := range products {
		if val.Product.Name == productName {
			return val.Prices
		}
	}
	return nil
}

var products = initializeProducts()

func initializeProducts() []*pb.CompareProductList {
	// Using make to create a slice of CompareProductList structs with initial values
	products := make([]*pb.CompareProductList, 0)

	// Assigning initial values to the slice
	products = append(
		products,
		&pb.CompareProductList{
			Product: &pb.Product{
				Id:       "1",
				Name:     "Dougnuts",
				ImageUrl: "https://images.unsplash.com/photo-1597586255676-6b33b6268e33?q=80&w=2052&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
				Details:  "Delicious donuts freshly baked every morning",
			},
			Prices: []*pb.MarketPrice{
				{MarketName: "walmart", Price: 30},
			},
		})

	// Fake data
	// TODO: replace with data from the database or any other source
	additionalProducts := []*pb.CompareProductList{
		{
			Product: &pb.Product{Id: "2", Name: "Bagels", ImageUrl: "https://images.unsplash.com/photo-1664038082440-ee0a677af315?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D", Details: "Freshly baked bagels perfect for breakfast"},
			Prices: []*pb.MarketPrice{
				{MarketName: "costco", Price: 25},
				{MarketName: "local grocery", Price: 27},
			},
		},
		{
			Product: &pb.Product{Id: "3", Name: "Croissants", ImageUrl: "https://images.unsplash.com/photo-1623334044303-241021148842?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D", Details: "Flaky and buttery croissants made with the finest ingredients"},
			Prices: []*pb.MarketPrice{
				{MarketName: "bakery", Price: 35},
				{MarketName: "cafe", Price: 38},
			},
		},
		// Add more products with their prices similarly
	}
	products = append(products, additionalProducts...)
	return products
}
