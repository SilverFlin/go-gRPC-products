package model

import (
	"context"
	"fmt"
	"log"
	"slices"

	"github.com/silverflin/go-rpc/internal/database"
	pb "github.com/silverflin/go-rpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var products []*pb.CompareProductList

func init() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer database.Close()
	products = initializeProducts()
}

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

func initializeProducts() []*pb.CompareProductList {

	collection := database.Client.Database("profeco-products").Collection("CompareProductList")

	collection.DeleteMany(context.Background(), bson.M{})
	products := []*pb.CompareProductList{
		{
			Product: &pb.Product{
				Name:     "Dougnuts",
				ImageUrl: "https://images.unsplash.com/photo-1597586255676-6b33b6268e33?q=80&w=2052&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fH",
				Details:  "Delicious donuts freshly baked every morning",
			},
			Prices: []*pb.MarketPrice{
				{MarketName: "walmart", Price: 30},
			},
		},

		{
			Product: &pb.Product{
				Name:     "Bagels",
				ImageUrl: "https://images.unsplash.com/photo-1664038082440-ee0a677af315?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M",
			},
			Prices: []*pb.MarketPrice{
				{MarketName: "costco", Price: 25},
				{MarketName: "local grocery", Price: 27},
			},
		},
		{
			Product: &pb.Product{
				Name:     "Croissants",
				ImageUrl: "https://images.unsplash.com/photo-1623334044303-241021148842?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=",
			},
			Prices: []*pb.MarketPrice{
				{MarketName: "bakery", Price: 35},
				{MarketName: "cafe", Price: 38},
			},
		},
	}

	for _, product := range products {
		result, err := collection.InsertOne(context.Background(), product)
		if err != nil {
			log.Fatal(err)
		}

		insertedID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			log.Fatal("Failed to convert inserted ID to ObjectID")
		}

		fmt.Println(insertedID.Hex())

		product.Product.Id = insertedID.Hex()
	}

	products = products[:0]

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product pb.CompareProductList

		if err := cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}

		var result map[string]interface{}

		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		id, ok := result["_id"].(primitive.ObjectID)
		if !ok {
			log.Fatal("Failed to convert _id to ObjectID")
		}

		product.Product.Id = id.Hex()
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return products
}
