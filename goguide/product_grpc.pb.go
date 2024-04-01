// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: goguide/product.proto

package go_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsClient interface {
	GetProductsByPrice(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductList, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetProductsByPrice(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/goguide.Products/GetProductsByPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
// All implementations must embed UnimplementedProductsServer
// for forward compatibility
type ProductsServer interface {
	GetProductsByPrice(context.Context, *ProductListRequest) (*ProductList, error)
	mustEmbedUnimplementedProductsServer()
}

// UnimplementedProductsServer must be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (UnimplementedProductsServer) GetProductsByPrice(context.Context, *ProductListRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByPrice not implemented")
}
func (UnimplementedProductsServer) mustEmbedUnimplementedProductsServer() {}

// UnsafeProductsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServer will
// result in compilation errors.
type UnsafeProductsServer interface {
	mustEmbedUnimplementedProductsServer()
}

func RegisterProductsServer(s grpc.ServiceRegistrar, srv ProductsServer) {
	s.RegisterService(&Products_ServiceDesc, srv)
}

func _Products_GetProductsByPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductsByPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goguide.Products/GetProductsByPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductsByPrice(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Products_ServiceDesc is the grpc.ServiceDesc for Products service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Products_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goguide.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductsByPrice",
			Handler:    _Products_GetProductsByPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goguide/product.proto",
}
