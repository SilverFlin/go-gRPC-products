package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/silverflin/go-rpc/internal/services"
	pb "github.com/silverflin/go-rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	var err error

	if os.Getenv("PRODUCT_MICROSERVICE_ENV") == "PROD" {
		err = godotenv.Load(".env.production")
	} else {
		err = godotenv.Load(".env.development")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	rpcURI := fmt.Sprintf("%v:%v", os.Getenv("RPC_URL"), os.Getenv("RPC_PORT"))
	lis, err := net.Listen("tcp", rpcURI)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Print(fmt.Sprintf("Listening on %v", rpcURI))

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductsServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *services.ProductListServer {
	s := &services.ProductListServer{}
	return s
}
