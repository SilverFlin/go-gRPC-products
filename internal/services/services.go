package services

import (
	"context"
	"github.com/silverflin/go-rpc/internal/messaging"
	"github.com/silverflin/go-rpc/internal/model"
	pb "github.com/silverflin/go-rpc/proto"
	"log"
	"sort"
)

type ProductListServer struct {
	pb.UnimplementedProductsServer
}

func (s ProductListServer) GetProductsByPrice(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductList, error) {
	log.Printf("New Request %v", req.ProductName)
	filteredProductList := &pb.ProductList{Products: make([]*pb.Product, 0)}
	for _, prod := range model.GetAllProducts() {
		if prod.Name == req.ProductName {
			filteredProductList.Products = append(filteredProductList.Products, prod)
		}
	}
	OrderProductListByPrice(filteredProductList)

	go messaging.SendToProductQueue("List request")

	return filteredProductList, nil
}

func (s ProductListServer) GetProductsNames(ctx context.Context, empty *pb.Empty) (*pb.ProductNamesList, error) {
	log.Print("New Request: products names")

	productsNames := model.GetProductsNames()

	go messaging.SendToProductQueue("Products Names Request")
	return productsNames, nil
}

func OrderProductListByPrice(l *pb.ProductList) {
	sort.Slice(l.Products, func(i, j int) bool {
		return l.Products[i].CurrentPrice > l.Products[j].CurrentPrice
	})
}
