package main

import (
	"context"
	"fmt"
	pb "github.com/silverflin/go-rpc/goguide"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var opts []grpc.DialOption
	//grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductsClient(conn)

	productList, err := client.GetProductsByPrice(context.Background(), &pb.ProductListRequest{ProductName: "Bagels"})
	if err != nil {
		fmt.Println("error at getting products")
	}

	log.Println(productList)

}

//
