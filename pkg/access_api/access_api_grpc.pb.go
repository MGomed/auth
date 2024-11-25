// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: access_api.proto

package access_api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccessAPIClient is the client API for AccessAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessAPIClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type accessAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessAPIClient(cc grpc.ClientConnInterface) AccessAPIClient {
	return &accessAPIClient{cc}
}

func (c *accessAPIClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access_api.AccessAPI/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessAPIServer is the server API for AccessAPI service.
// All implementations must embed UnimplementedAccessAPIServer
// for forward compatibility
type AccessAPIServer interface {
	Check(context.Context, *CheckRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAccessAPIServer()
}

// UnimplementedAccessAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAccessAPIServer struct {
}

func (UnimplementedAccessAPIServer) Check(context.Context, *CheckRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAccessAPIServer) mustEmbedUnimplementedAccessAPIServer() {}

// UnsafeAccessAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessAPIServer will
// result in compilation errors.
type UnsafeAccessAPIServer interface {
	mustEmbedUnimplementedAccessAPIServer()
}

func RegisterAccessAPIServer(s grpc.ServiceRegistrar, srv AccessAPIServer) {
	s.RegisterService(&AccessAPI_ServiceDesc, srv)
}

func _AccessAPI_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access_api.AccessAPI/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessAPI_ServiceDesc is the grpc.ServiceDesc for AccessAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "access_api.AccessAPI",
	HandlerType: (*AccessAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _AccessAPI_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access_api.proto",
}
