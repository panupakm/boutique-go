// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: catalog/catalog.proto

package catalog

import (
	context "context"
	shared "github.com/panupakm/boutique-go/api/shared"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Catalog_ListProducts_FullMethodName   = "/boutiqueshop.Catalog/ListProducts"
	Catalog_GetProduct_FullMethodName     = "/boutiqueshop.Catalog/GetProduct"
	Catalog_SearchProducts_FullMethodName = "/boutiqueshop.Catalog/SearchProducts"
)

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*shared.Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, Catalog_ListProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*shared.Product, error) {
	out := new(shared.Product)
	err := c.cc.Invoke(ctx, Catalog_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, Catalog_SearchProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*shared.Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedCatalogServer) GetProduct(context.Context, *GetProductRequest) (*shared.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedCatalogServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_SearchProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "boutiqueshop.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _Catalog_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Catalog_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _Catalog_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/catalog.proto",
}
