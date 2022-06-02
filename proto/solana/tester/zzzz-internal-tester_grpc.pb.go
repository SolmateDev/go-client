// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: zzzz-internal-tester.proto

package tester

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

// ValidatorEnvClient is the client API for ValidatorEnv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidatorEnvClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (ValidatorEnv_CreateClient, error)
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error)
}

type validatorEnvClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorEnvClient(cc grpc.ClientConnInterface) ValidatorEnvClient {
	return &validatorEnvClient{cc}
}

func (c *validatorEnvClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (ValidatorEnv_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &ValidatorEnv_ServiceDesc.Streams[0], "/tester.ValidatorEnv/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &validatorEnvCreateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ValidatorEnv_CreateClient interface {
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type validatorEnvCreateClient struct {
	grpc.ClientStream
}

func (x *validatorEnvCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *validatorEnvClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error) {
	out := new(DestroyResponse)
	err := c.cc.Invoke(ctx, "/tester.ValidatorEnv/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorEnvServer is the server API for ValidatorEnv service.
// All implementations must embed UnimplementedValidatorEnvServer
// for forward compatibility
type ValidatorEnvServer interface {
	Create(*CreateRequest, ValidatorEnv_CreateServer) error
	Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error)
	mustEmbedUnimplementedValidatorEnvServer()
}

// UnimplementedValidatorEnvServer must be embedded to have forward compatible implementations.
type UnimplementedValidatorEnvServer struct {
}

func (UnimplementedValidatorEnvServer) Create(*CreateRequest, ValidatorEnv_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedValidatorEnvServer) Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedValidatorEnvServer) mustEmbedUnimplementedValidatorEnvServer() {}

// UnsafeValidatorEnvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidatorEnvServer will
// result in compilation errors.
type UnsafeValidatorEnvServer interface {
	mustEmbedUnimplementedValidatorEnvServer()
}

func RegisterValidatorEnvServer(s grpc.ServiceRegistrar, srv ValidatorEnvServer) {
	s.RegisterService(&ValidatorEnv_ServiceDesc, srv)
}

func _ValidatorEnv_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ValidatorEnvServer).Create(m, &validatorEnvCreateServer{stream})
}

type ValidatorEnv_CreateServer interface {
	Send(*CreateResponse) error
	grpc.ServerStream
}

type validatorEnvCreateServer struct {
	grpc.ServerStream
}

func (x *validatorEnvCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ValidatorEnv_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorEnvServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tester.ValidatorEnv/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorEnvServer).Destroy(ctx, req.(*DestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidatorEnv_ServiceDesc is the grpc.ServiceDesc for ValidatorEnv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidatorEnv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tester.ValidatorEnv",
	HandlerType: (*ValidatorEnvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Destroy",
			Handler:    _ValidatorEnv_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _ValidatorEnv_Create_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zzzz-internal-tester.proto",
}