// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: crudchain/crudchain/query.proto

package crudchain

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

const (
	Query_Params_FullMethodName             = "/crudchain.crudchain.Query/Params"
	Query_Resource_FullMethodName           = "/crudchain.crudchain.Query/Resource"
	Query_ResourceAll_FullMethodName        = "/crudchain.crudchain.Query/ResourceAll"
	Query_ShowResourceByName_FullMethodName = "/crudchain.crudchain.Query/ShowResourceByName"
	Query_Film_FullMethodName               = "/crudchain.crudchain.Query/Film"
	Query_FilmAll_FullMethodName            = "/crudchain.crudchain.Query/FilmAll"
	Query_ShowFilmByGenre_FullMethodName    = "/crudchain.crudchain.Query/ShowFilmByGenre"
	Query_ShowFilmGenre_FullMethodName      = "/crudchain.crudchain.Query/ShowFilmGenre"
	Query_ShowFilmName_FullMethodName       = "/crudchain.crudchain.Query/ShowFilmName"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Resource items.
	Resource(ctx context.Context, in *QueryGetResourceRequest, opts ...grpc.CallOption) (*QueryGetResourceResponse, error)
	ResourceAll(ctx context.Context, in *QueryAllResourceRequest, opts ...grpc.CallOption) (*QueryAllResourceResponse, error)
	// Queries a list of ShowResourceByName items.
	ShowResourceByName(ctx context.Context, in *QueryShowResourceByNameRequest, opts ...grpc.CallOption) (*QueryShowResourceByNameResponse, error)
	// Queries a list of Film items.
	Film(ctx context.Context, in *QueryGetFilmRequest, opts ...grpc.CallOption) (*QueryGetFilmResponse, error)
	FilmAll(ctx context.Context, in *QueryAllFilmRequest, opts ...grpc.CallOption) (*QueryAllFilmResponse, error)
	// Queries a list of ShowFilmByGenre items.
	ShowFilmByGenre(ctx context.Context, in *QueryShowFilmByGenreRequest, opts ...grpc.CallOption) (*QueryShowFilmByGenreResponse, error)
	// Queries a list of ShowFilmGenre items.
	ShowFilmGenre(ctx context.Context, in *QueryShowFilmGenreRequest, opts ...grpc.CallOption) (*QueryShowFilmGenreResponse, error)
	// Queries a list of ShowFilmName items.
	ShowFilmName(ctx context.Context, in *QueryShowFilmNameRequest, opts ...grpc.CallOption) (*QueryShowFilmNameResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Resource(ctx context.Context, in *QueryGetResourceRequest, opts ...grpc.CallOption) (*QueryGetResourceResponse, error) {
	out := new(QueryGetResourceResponse)
	err := c.cc.Invoke(ctx, Query_Resource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResourceAll(ctx context.Context, in *QueryAllResourceRequest, opts ...grpc.CallOption) (*QueryAllResourceResponse, error) {
	out := new(QueryAllResourceResponse)
	err := c.cc.Invoke(ctx, Query_ResourceAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowResourceByName(ctx context.Context, in *QueryShowResourceByNameRequest, opts ...grpc.CallOption) (*QueryShowResourceByNameResponse, error) {
	out := new(QueryShowResourceByNameResponse)
	err := c.cc.Invoke(ctx, Query_ShowResourceByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Film(ctx context.Context, in *QueryGetFilmRequest, opts ...grpc.CallOption) (*QueryGetFilmResponse, error) {
	out := new(QueryGetFilmResponse)
	err := c.cc.Invoke(ctx, Query_Film_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FilmAll(ctx context.Context, in *QueryAllFilmRequest, opts ...grpc.CallOption) (*QueryAllFilmResponse, error) {
	out := new(QueryAllFilmResponse)
	err := c.cc.Invoke(ctx, Query_FilmAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowFilmByGenre(ctx context.Context, in *QueryShowFilmByGenreRequest, opts ...grpc.CallOption) (*QueryShowFilmByGenreResponse, error) {
	out := new(QueryShowFilmByGenreResponse)
	err := c.cc.Invoke(ctx, Query_ShowFilmByGenre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowFilmGenre(ctx context.Context, in *QueryShowFilmGenreRequest, opts ...grpc.CallOption) (*QueryShowFilmGenreResponse, error) {
	out := new(QueryShowFilmGenreResponse)
	err := c.cc.Invoke(ctx, Query_ShowFilmGenre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowFilmName(ctx context.Context, in *QueryShowFilmNameRequest, opts ...grpc.CallOption) (*QueryShowFilmNameResponse, error) {
	out := new(QueryShowFilmNameResponse)
	err := c.cc.Invoke(ctx, Query_ShowFilmName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Resource items.
	Resource(context.Context, *QueryGetResourceRequest) (*QueryGetResourceResponse, error)
	ResourceAll(context.Context, *QueryAllResourceRequest) (*QueryAllResourceResponse, error)
	// Queries a list of ShowResourceByName items.
	ShowResourceByName(context.Context, *QueryShowResourceByNameRequest) (*QueryShowResourceByNameResponse, error)
	// Queries a list of Film items.
	Film(context.Context, *QueryGetFilmRequest) (*QueryGetFilmResponse, error)
	FilmAll(context.Context, *QueryAllFilmRequest) (*QueryAllFilmResponse, error)
	// Queries a list of ShowFilmByGenre items.
	ShowFilmByGenre(context.Context, *QueryShowFilmByGenreRequest) (*QueryShowFilmByGenreResponse, error)
	// Queries a list of ShowFilmGenre items.
	ShowFilmGenre(context.Context, *QueryShowFilmGenreRequest) (*QueryShowFilmGenreResponse, error)
	// Queries a list of ShowFilmName items.
	ShowFilmName(context.Context, *QueryShowFilmNameRequest) (*QueryShowFilmNameResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Resource(context.Context, *QueryGetResourceRequest) (*QueryGetResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resource not implemented")
}
func (UnimplementedQueryServer) ResourceAll(context.Context, *QueryAllResourceRequest) (*QueryAllResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceAll not implemented")
}
func (UnimplementedQueryServer) ShowResourceByName(context.Context, *QueryShowResourceByNameRequest) (*QueryShowResourceByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowResourceByName not implemented")
}
func (UnimplementedQueryServer) Film(context.Context, *QueryGetFilmRequest) (*QueryGetFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Film not implemented")
}
func (UnimplementedQueryServer) FilmAll(context.Context, *QueryAllFilmRequest) (*QueryAllFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilmAll not implemented")
}
func (UnimplementedQueryServer) ShowFilmByGenre(context.Context, *QueryShowFilmByGenreRequest) (*QueryShowFilmByGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowFilmByGenre not implemented")
}
func (UnimplementedQueryServer) ShowFilmGenre(context.Context, *QueryShowFilmGenreRequest) (*QueryShowFilmGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowFilmGenre not implemented")
}
func (UnimplementedQueryServer) ShowFilmName(context.Context, *QueryShowFilmNameRequest) (*QueryShowFilmNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowFilmName not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Resource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Resource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Resource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Resource(ctx, req.(*QueryGetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResourceAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResourceAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ResourceAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResourceAll(ctx, req.(*QueryAllResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowResourceByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowResourceByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowResourceByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowResourceByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowResourceByName(ctx, req.(*QueryShowResourceByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Film_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Film(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Film_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Film(ctx, req.(*QueryGetFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FilmAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FilmAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_FilmAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FilmAll(ctx, req.(*QueryAllFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowFilmByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowFilmByGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowFilmByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowFilmByGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowFilmByGenre(ctx, req.(*QueryShowFilmByGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowFilmGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowFilmGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowFilmGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowFilmGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowFilmGenre(ctx, req.(*QueryShowFilmGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowFilmName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowFilmNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowFilmName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowFilmName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowFilmName(ctx, req.(*QueryShowFilmNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crudchain.crudchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Resource",
			Handler:    _Query_Resource_Handler,
		},
		{
			MethodName: "ResourceAll",
			Handler:    _Query_ResourceAll_Handler,
		},
		{
			MethodName: "ShowResourceByName",
			Handler:    _Query_ShowResourceByName_Handler,
		},
		{
			MethodName: "Film",
			Handler:    _Query_Film_Handler,
		},
		{
			MethodName: "FilmAll",
			Handler:    _Query_FilmAll_Handler,
		},
		{
			MethodName: "ShowFilmByGenre",
			Handler:    _Query_ShowFilmByGenre_Handler,
		},
		{
			MethodName: "ShowFilmGenre",
			Handler:    _Query_ShowFilmGenre_Handler,
		},
		{
			MethodName: "ShowFilmName",
			Handler:    _Query_ShowFilmName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crudchain/crudchain/query.proto",
}
