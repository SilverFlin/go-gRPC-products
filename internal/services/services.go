package services

import (
	"context"
	pb "github.com/silverflin/go-rpc/proto"
	"github.com/silverflin/go-rpc/internal/messaging"
	"github.com/silverflin/go-rpc/internal/model"
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

	messaging.SendToProductQueue("List request")

	return filteredProductList, nil
}

func OrderProductListByPrice(l *pb.ProductList) {
	sort.Slice(l.Products, func(i, j int) bool {
		return l.Products[i].CurrentPrice > l.Products[j].CurrentPrice
	})
}
