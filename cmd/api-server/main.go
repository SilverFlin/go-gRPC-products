package main

import (
	"fmt"
	pb "github.com/silverflin/go-rpc/proto"
	"github.com/silverflin/go-rpc/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

// add prod/build vari
var URL = ""
var PORT = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", URL, PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Print(fmt.Sprintf("Listening on %v:%d", URL, PORT))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductsServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *services.ProductListServer {
	s := &services.ProductListServer{}
	return s
}
