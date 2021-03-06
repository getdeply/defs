// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

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

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error)
	GetBuildLog(ctx context.Context, in *GetBuildLogRequest, opts ...grpc.CallOption) (Services_GetBuildLogClient, error)
	ListGithubRepos(ctx context.Context, in *ListGithubReposRequest, opts ...grpc.CallOption) (*ListGithubReposResponse, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.services.Services/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.services.Services/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/api.services.Services/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error) {
	out := new(ListVersionsResponse)
	err := c.cc.Invoke(ctx, "/api.services.Services/ListVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetBuildLog(ctx context.Context, in *GetBuildLogRequest, opts ...grpc.CallOption) (Services_GetBuildLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &Services_ServiceDesc.Streams[0], "/api.services.Services/GetBuildLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesGetBuildLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Services_GetBuildLogClient interface {
	Recv() (*GetBuildLogResponse, error)
	grpc.ClientStream
}

type servicesGetBuildLogClient struct {
	grpc.ClientStream
}

func (x *servicesGetBuildLogClient) Recv() (*GetBuildLogResponse, error) {
	m := new(GetBuildLogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicesClient) ListGithubRepos(ctx context.Context, in *ListGithubReposRequest, opts ...grpc.CallOption) (*ListGithubReposResponse, error) {
	out := new(ListGithubReposResponse)
	err := c.cc.Invoke(ctx, "/api.services.Services/ListGithubRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
// All implementations must embed UnimplementedServicesServer
// for forward compatibility
type ServicesServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error)
	GetBuildLog(*GetBuildLogRequest, Services_GetBuildLogServer) error
	ListGithubRepos(context.Context, *ListGithubReposRequest) (*ListGithubReposResponse, error)
	mustEmbedUnimplementedServicesServer()
}

// UnimplementedServicesServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServer struct {
}

func (UnimplementedServicesServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServicesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServicesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServicesServer) ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}
func (UnimplementedServicesServer) GetBuildLog(*GetBuildLogRequest, Services_GetBuildLogServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBuildLog not implemented")
}
func (UnimplementedServicesServer) ListGithubRepos(context.Context, *ListGithubReposRequest) (*ListGithubReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGithubRepos not implemented")
}
func (UnimplementedServicesServer) mustEmbedUnimplementedServicesServer() {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.Services/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.Services/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.Services/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.Services/ListVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ListVersions(ctx, req.(*ListVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetBuildLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBuildLogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicesServer).GetBuildLog(m, &servicesGetBuildLogServer{stream})
}

type Services_GetBuildLogServer interface {
	Send(*GetBuildLogResponse) error
	grpc.ServerStream
}

type servicesGetBuildLogServer struct {
	grpc.ServerStream
}

func (x *servicesGetBuildLogServer) Send(m *GetBuildLogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Services_ListGithubRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGithubReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ListGithubRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.Services/ListGithubRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ListGithubRepos(ctx, req.(*ListGithubReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Services_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Services_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Services_List_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _Services_ListVersions_Handler,
		},
		{
			MethodName: "ListGithubRepos",
			Handler:    _Services_ListGithubRepos_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBuildLog",
			Handler:       _Services_GetBuildLog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/services/services.proto",
}
