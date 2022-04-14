// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: mds.proto

package mds

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

// MdsServiceClient is the client API for MdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MdsServiceClient interface {
	ListClusterFsInfo(ctx context.Context, in *ListClusterFsInfoRequest, opts ...grpc.CallOption) (*ListClusterFsInfoResponse, error)
	DeleteFs(ctx context.Context, in *DeleteFsRequest, opts ...grpc.CallOption) (*DeleteFsResponse, error)
}

type mdsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMdsServiceClient(cc grpc.ClientConnInterface) MdsServiceClient {
	return &mdsServiceClient{cc}
}

func (c *mdsServiceClient) ListClusterFsInfo(ctx context.Context, in *ListClusterFsInfoRequest, opts ...grpc.CallOption) (*ListClusterFsInfoResponse, error) {
	out := new(ListClusterFsInfoResponse)
	err := c.cc.Invoke(ctx, "/curvefs.mds.MdsService/ListClusterFsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mdsServiceClient) DeleteFs(ctx context.Context, in *DeleteFsRequest, opts ...grpc.CallOption) (*DeleteFsResponse, error) {
	out := new(DeleteFsResponse)
	err := c.cc.Invoke(ctx, "/curvefs.mds.MdsService/DeleteFs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MdsServiceServer is the server API for MdsService service.
// All implementations must embed UnimplementedMdsServiceServer
// for forward compatibility
type MdsServiceServer interface {
	ListClusterFsInfo(context.Context, *ListClusterFsInfoRequest) (*ListClusterFsInfoResponse, error)
	DeleteFs(context.Context, *DeleteFsRequest) (*DeleteFsResponse, error)
	mustEmbedUnimplementedMdsServiceServer()
}

// UnimplementedMdsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMdsServiceServer struct {
}

func (UnimplementedMdsServiceServer) ListClusterFsInfo(context.Context, *ListClusterFsInfoRequest) (*ListClusterFsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusterFsInfo not implemented")
}
func (UnimplementedMdsServiceServer) DeleteFs(context.Context, *DeleteFsRequest) (*DeleteFsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFs not implemented")
}
func (UnimplementedMdsServiceServer) mustEmbedUnimplementedMdsServiceServer() {}

// UnsafeMdsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MdsServiceServer will
// result in compilation errors.
type UnsafeMdsServiceServer interface {
	mustEmbedUnimplementedMdsServiceServer()
}

func RegisterMdsServiceServer(s grpc.ServiceRegistrar, srv MdsServiceServer) {
	s.RegisterService(&MdsService_ServiceDesc, srv)
}

func _MdsService_ListClusterFsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterFsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MdsServiceServer).ListClusterFsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/curvefs.mds.MdsService/ListClusterFsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MdsServiceServer).ListClusterFsInfo(ctx, req.(*ListClusterFsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MdsService_DeleteFs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MdsServiceServer).DeleteFs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/curvefs.mds.MdsService/DeleteFs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MdsServiceServer).DeleteFs(ctx, req.(*DeleteFsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MdsService_ServiceDesc is the grpc.ServiceDesc for MdsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MdsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "curvefs.mds.MdsService",
	HandlerType: (*MdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClusterFsInfo",
			Handler:    _MdsService_ListClusterFsInfo_Handler,
		},
		{
			MethodName: "DeleteFs",
			Handler:    _MdsService_DeleteFs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mds.proto",
}