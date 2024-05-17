// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DirectorCommunicationClient is the client API for DirectorCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectorCommunicationClient interface {
	ConfirmReady(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	PickGun(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	PickHallway(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	BossBattle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	AskForAccountBalance(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type directorCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectorCommunicationClient(cc grpc.ClientConnInterface) DirectorCommunicationClient {
	return &directorCommunicationClient{cc}
}

func (c *directorCommunicationClient) ConfirmReady(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.DirectorCommunication/ConfirmReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorCommunicationClient) PickGun(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.DirectorCommunication/PickGun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorCommunicationClient) PickHallway(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.DirectorCommunication/PickHallway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorCommunicationClient) BossBattle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.DirectorCommunication/BossBattle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorCommunicationClient) AskForAccountBalance(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.DirectorCommunication/AskForAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectorCommunicationServer is the server API for DirectorCommunication service.
// All implementations must embed UnimplementedDirectorCommunicationServer
// for forward compatibility
type DirectorCommunicationServer interface {
	ConfirmReady(context.Context, *Message) (*Response, error)
	PickGun(context.Context, *Message) (*Response, error)
	PickHallway(context.Context, *Message) (*Response, error)
	BossBattle(context.Context, *Message) (*Response, error)
	AskForAccountBalance(context.Context, *Message) (*Response, error)
	mustEmbedUnimplementedDirectorCommunicationServer()
}

// UnimplementedDirectorCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedDirectorCommunicationServer struct {
}

func (UnimplementedDirectorCommunicationServer) ConfirmReady(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReady not implemented")
}
func (UnimplementedDirectorCommunicationServer) PickGun(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickGun not implemented")
}
func (UnimplementedDirectorCommunicationServer) PickHallway(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickHallway not implemented")
}
func (UnimplementedDirectorCommunicationServer) BossBattle(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BossBattle not implemented")
}
func (UnimplementedDirectorCommunicationServer) AskForAccountBalance(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForAccountBalance not implemented")
}
func (UnimplementedDirectorCommunicationServer) mustEmbedUnimplementedDirectorCommunicationServer() {}

// UnsafeDirectorCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectorCommunicationServer will
// result in compilation errors.
type UnsafeDirectorCommunicationServer interface {
	mustEmbedUnimplementedDirectorCommunicationServer()
}

func RegisterDirectorCommunicationServer(s *grpc.Server, srv DirectorCommunicationServer) {
	s.RegisterService(&_DirectorCommunication_serviceDesc, srv)
}

func _DirectorCommunication_ConfirmReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorCommunicationServer).ConfirmReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.DirectorCommunication/ConfirmReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorCommunicationServer).ConfirmReady(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectorCommunication_PickGun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorCommunicationServer).PickGun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.DirectorCommunication/PickGun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorCommunicationServer).PickGun(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectorCommunication_PickHallway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorCommunicationServer).PickHallway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.DirectorCommunication/PickHallway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorCommunicationServer).PickHallway(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectorCommunication_BossBattle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorCommunicationServer).BossBattle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.DirectorCommunication/BossBattle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorCommunicationServer).BossBattle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectorCommunication_AskForAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorCommunicationServer).AskForAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.DirectorCommunication/AskForAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorCommunicationServer).AskForAccountBalance(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _DirectorCommunication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "director.DirectorCommunication",
	HandlerType: (*DirectorCommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfirmReady",
			Handler:    _DirectorCommunication_ConfirmReady_Handler,
		},
		{
			MethodName: "PickGun",
			Handler:    _DirectorCommunication_PickGun_Handler,
		},
		{
			MethodName: "PickHallway",
			Handler:    _DirectorCommunication_PickHallway_Handler,
		},
		{
			MethodName: "BossBattle",
			Handler:    _DirectorCommunication_BossBattle_Handler,
		},
		{
			MethodName: "AskForAccountBalance",
			Handler:    _DirectorCommunication_AskForAccountBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/director.proto",
}

// NodesCommunicationClient is the client API for NodesCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodesCommunicationClient interface {
	NotifyDecision(ctx context.Context, in *Decision, opts ...grpc.CallOption) (*Response, error)
}

type nodesCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewNodesCommunicationClient(cc grpc.ClientConnInterface) NodesCommunicationClient {
	return &nodesCommunicationClient{cc}
}

func (c *nodesCommunicationClient) NotifyDecision(ctx context.Context, in *Decision, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/director.NodesCommunication/NotifyDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesCommunicationServer is the server API for NodesCommunication service.
// All implementations must embed UnimplementedNodesCommunicationServer
// for forward compatibility
type NodesCommunicationServer interface {
	NotifyDecision(context.Context, *Decision) (*Response, error)
	mustEmbedUnimplementedNodesCommunicationServer()
}

// UnimplementedNodesCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedNodesCommunicationServer struct {
}

func (UnimplementedNodesCommunicationServer) NotifyDecision(context.Context, *Decision) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyDecision not implemented")
}
func (UnimplementedNodesCommunicationServer) mustEmbedUnimplementedNodesCommunicationServer() {}

// UnsafeNodesCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodesCommunicationServer will
// result in compilation errors.
type UnsafeNodesCommunicationServer interface {
	mustEmbedUnimplementedNodesCommunicationServer()
}

func RegisterNodesCommunicationServer(s *grpc.Server, srv NodesCommunicationServer) {
	s.RegisterService(&_NodesCommunication_serviceDesc, srv)
}

func _NodesCommunication_NotifyDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Decision)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesCommunicationServer).NotifyDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/director.NodesCommunication/NotifyDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesCommunicationServer).NotifyDecision(ctx, req.(*Decision))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodesCommunication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "director.NodesCommunication",
	HandlerType: (*NodesCommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyDecision",
			Handler:    _NodesCommunication_NotifyDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/director.proto",
}
