package model

import (
	pb "github.com/silverflin/go-rpc/proto"
	time "google.golang.org/protobuf/types/known/timestamppb"
)

func GetAllProducts() []*pb.Product {
	return products
}

var products = initializeProducts()

func initializeProducts() []*pb.Product {
	// Using make to create a slice of Product structs with initial values
	products := make([]*pb.Product, 0)

	// Assigning initial values to the slice
	products = append(
		products,
		&pb.Product{
			Id:              "1",
			Name:            "Dougnuts",
			CurrentPrice:    30,
			LastUpdatedTime: time.Now(),
			Market:          "walmart",
		})

	// fake data
	// TODO replace for db persistence data
	additionalProducts := []*pb.Product{
		{Id: "2", Name: "Bagels", CurrentPrice: 25, LastUpdatedTime: time.Now(), Market: "costco"},
		{Id: "3", Name: "Croissants", CurrentPrice: 35, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "4", Name: "Muffins", CurrentPrice: 20, LastUpdatedTime: time.Now(), Market: "grocery store"},
		{Id: "5", Name: "Danishes", CurrentPrice: 28, LastUpdatedTime: time.Now(), Market: "local bakery"},
		{Id: "6", Name: "Cinnamon Rolls", CurrentPrice: 32, LastUpdatedTime: time.Now(), Market: "cafe"},
		{Id: "7", Name: "Scones", CurrentPrice: 22, LastUpdatedTime: time.Now(), Market: "coffee shop"},
		{Id: "8", Name: "Biscuits", CurrentPrice: 18, LastUpdatedTime: time.Now(), Market: "restaurant"},
		{Id: "9", Name: "Pretzels", CurrentPrice: 15, LastUpdatedTime: time.Now(), Market: "street vendor"},
		{Id: "10", Name: "Cookies", CurrentPrice: 27, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "11", Name: "Bagels", CurrentPrice: 30, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "12", Name: "Cake", CurrentPrice: 40, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "13", Name: "Pie", CurrentPrice: 35, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "14", Name: "Cupcakes", CurrentPrice: 25, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "15", Name: "Cheesecake", CurrentPrice: 50, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "16", Name: "Tarts", CurrentPrice: 38, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "17", Name: "Eclairs", CurrentPrice: 45, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "18", Name: "Macarons", CurrentPrice: 55, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "19", Name: "Bagels", CurrentPrice: 10, LastUpdatedTime: time.Now(), Market: "bakery"},
		{Id: "20", Name: "Rolls", CurrentPrice: 12, LastUpdatedTime: time.Now(), Market: "bakery"},
	}
	products = append(products, additionalProducts...)
	return products
}
