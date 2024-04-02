package main

import (
	"context"
	"fmt"
	pb "github.com/silverflin/go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
    conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductsClient(conn)

	fmt.Println("[Getting Product List: Bagels]")
	productList, err := client.GetProductsByPrice(context.Background(), &pb.ProductListRequest{ProductName: "Bagels"})
	if err != nil {
		fmt.Println("error at getting products")
	}
	fmt.Println(productList)

	fmt.Println("[Getting Product Names]")
	productNames, err := client.GetProductsNames(context.Background(), &pb.Empty{})

	if err != nil {
		fmt.Println("error at getting products names")
	}
	fmt.Println(productNames)
}
