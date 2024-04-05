package services

import (
	"context"
	"github.com/silverflin/go-rpc/internal/messaging"
	"github.com/silverflin/go-rpc/internal/model"
	pb "github.com/silverflin/go-rpc/proto"
	"log"
)

type ProductListServer struct {
	pb.UnimplementedProductsServer
}

func (s ProductListServer) GetProductsByPrice(ctx context.Context, req *pb.ProductListRequest) (*pb.CompareProductList, error) {
	log.Printf("New Request %v", req.ProductName)
	filteredProductList := &pb.CompareProductList{Product: &pb.Product{}, Prices: make([]*pb.MarketPrice, 0)}

	allProducts := model.GetProducts()

	for _, val := range allProducts.Product {
		if val.Name == req.ProductName {
			filteredProductList.Product = val
			filteredProductList.Prices = model.GetPricesFromProduct(val.Name)
			break
		}
	}

	go messaging.SendToProductQueue("List request")

	return filteredProductList, nil
}

func (s ProductListServer) GetProducts(ctx context.Context, empty *pb.Empty) (*pb.ProductList, error) {
	log.Print("New Request: products ")

	productsNames := model.GetProducts()

	go messaging.SendToProductQueue("Products Names Request")
	return productsNames, nil
}

