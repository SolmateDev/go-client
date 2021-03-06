// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: zzzz-customer-price.proto

package customer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	price "github.com/SolmateDev/go-client/proto/price"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PricingClient is the client API for Pricing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricingClient interface {
	//rpc GetStream(base.ShortCommodity) returns (stream PriceUpdate) {}
	GetLastTrade(ctx context.Context, in *price.MarketRequest, opts ...grpc.CallOption) (Pricing_GetLastTradeClient, error)
	GetOrderBook(ctx context.Context, in *price.MarketRequest, opts ...grpc.CallOption) (Pricing_GetOrderBookClient, error)
}

type pricingClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingClient(cc grpc.ClientConnInterface) PricingClient {
	return &pricingClient{cc}
}

func (c *pricingClient) GetLastTrade(ctx context.Context, in *price.MarketRequest, opts ...grpc.CallOption) (Pricing_GetLastTradeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pricing_ServiceDesc.Streams[0], "/customer.Pricing/GetLastTrade", opts...)
	if err != nil {
		return nil, err
	}
	x := &pricingGetLastTradeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pricing_GetLastTradeClient interface {
	Recv() (*price.Trade, error)
	grpc.ClientStream
}

type pricingGetLastTradeClient struct {
	grpc.ClientStream
}

func (x *pricingGetLastTradeClient) Recv() (*price.Trade, error) {
	m := new(price.Trade)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pricingClient) GetOrderBook(ctx context.Context, in *price.MarketRequest, opts ...grpc.CallOption) (Pricing_GetOrderBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pricing_ServiceDesc.Streams[1], "/customer.Pricing/GetOrderBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &pricingGetOrderBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pricing_GetOrderBookClient interface {
	Recv() (*price.OrderBook, error)
	grpc.ClientStream
}

type pricingGetOrderBookClient struct {
	grpc.ClientStream
}

func (x *pricingGetOrderBookClient) Recv() (*price.OrderBook, error) {
	m := new(price.OrderBook)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PricingServer is the server API for Pricing service.
// All implementations must embed UnimplementedPricingServer
// for forward compatibility
type PricingServer interface {
	//rpc GetStream(base.ShortCommodity) returns (stream PriceUpdate) {}
	GetLastTrade(*price.MarketRequest, Pricing_GetLastTradeServer) error
	GetOrderBook(*price.MarketRequest, Pricing_GetOrderBookServer) error
	mustEmbedUnimplementedPricingServer()
}

// UnimplementedPricingServer must be embedded to have forward compatible implementations.
type UnimplementedPricingServer struct {
}

func (UnimplementedPricingServer) GetLastTrade(*price.MarketRequest, Pricing_GetLastTradeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLastTrade not implemented")
}
func (UnimplementedPricingServer) GetOrderBook(*price.MarketRequest, Pricing_GetOrderBookServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderBook not implemented")
}
func (UnimplementedPricingServer) mustEmbedUnimplementedPricingServer() {}

// UnsafePricingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricingServer will
// result in compilation errors.
type UnsafePricingServer interface {
	mustEmbedUnimplementedPricingServer()
}

func RegisterPricingServer(s grpc.ServiceRegistrar, srv PricingServer) {
	s.RegisterService(&Pricing_ServiceDesc, srv)
}

func _Pricing_GetLastTrade_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(price.MarketRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PricingServer).GetLastTrade(m, &pricingGetLastTradeServer{stream})
}

type Pricing_GetLastTradeServer interface {
	Send(*price.Trade) error
	grpc.ServerStream
}

type pricingGetLastTradeServer struct {
	grpc.ServerStream
}

func (x *pricingGetLastTradeServer) Send(m *price.Trade) error {
	return x.ServerStream.SendMsg(m)
}

func _Pricing_GetOrderBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(price.MarketRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PricingServer).GetOrderBook(m, &pricingGetOrderBookServer{stream})
}

type Pricing_GetOrderBookServer interface {
	Send(*price.OrderBook) error
	grpc.ServerStream
}

type pricingGetOrderBookServer struct {
	grpc.ServerStream
}

func (x *pricingGetOrderBookServer) Send(m *price.OrderBook) error {
	return x.ServerStream.SendMsg(m)
}

// Pricing_ServiceDesc is the grpc.ServiceDesc for Pricing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pricing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Pricing",
	HandlerType: (*PricingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLastTrade",
			Handler:       _Pricing_GetLastTrade_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrderBook",
			Handler:       _Pricing_GetOrderBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zzzz-customer-price.proto",
}
