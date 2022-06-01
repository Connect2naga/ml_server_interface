// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: ml_server_interface.proto

package experiment

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

// IterationHandlerClient is the client API for IterationHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IterationHandlerClient interface {
	StartIterationExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*Null, error)
	StatusIterationExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*IterationResp, error)
}

type iterationHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewIterationHandlerClient(cc grpc.ClientConnInterface) IterationHandlerClient {
	return &iterationHandlerClient{cc}
}

func (c *iterationHandlerClient) StartIterationExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/experiment.IterationHandler/StartIterationExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iterationHandlerClient) StatusIterationExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*IterationResp, error) {
	out := new(IterationResp)
	err := c.cc.Invoke(ctx, "/experiment.IterationHandler/StatusIterationExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IterationHandlerServer is the server API for IterationHandler service.
// All implementations must embed UnimplementedIterationHandlerServer
// for forward compatibility
type IterationHandlerServer interface {
	StartIterationExperiment(context.Context, *Experiment) (*Null, error)
	StatusIterationExperiment(context.Context, *Experiment) (*IterationResp, error)
	mustEmbedUnimplementedIterationHandlerServer()
}

// UnimplementedIterationHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedIterationHandlerServer struct {
}

func (UnimplementedIterationHandlerServer) StartIterationExperiment(context.Context, *Experiment) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartIterationExperiment not implemented")
}
func (UnimplementedIterationHandlerServer) StatusIterationExperiment(context.Context, *Experiment) (*IterationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusIterationExperiment not implemented")
}
func (UnimplementedIterationHandlerServer) mustEmbedUnimplementedIterationHandlerServer() {}

// UnsafeIterationHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IterationHandlerServer will
// result in compilation errors.
type UnsafeIterationHandlerServer interface {
	mustEmbedUnimplementedIterationHandlerServer()
}

func RegisterIterationHandlerServer(s grpc.ServiceRegistrar, srv IterationHandlerServer) {
	s.RegisterService(&IterationHandler_ServiceDesc, srv)
}

func _IterationHandler_StartIterationExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Experiment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IterationHandlerServer).StartIterationExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.IterationHandler/StartIterationExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IterationHandlerServer).StartIterationExperiment(ctx, req.(*Experiment))
	}
	return interceptor(ctx, in, info, handler)
}

func _IterationHandler_StatusIterationExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Experiment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IterationHandlerServer).StatusIterationExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.IterationHandler/StatusIterationExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IterationHandlerServer).StatusIterationExperiment(ctx, req.(*Experiment))
	}
	return interceptor(ctx, in, info, handler)
}

// IterationHandler_ServiceDesc is the grpc.ServiceDesc for IterationHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IterationHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiment.IterationHandler",
	HandlerType: (*IterationHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartIterationExperiment",
			Handler:    _IterationHandler_StartIterationExperiment_Handler,
		},
		{
			MethodName: "StatusIterationExperiment",
			Handler:    _IterationHandler_StatusIterationExperiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ml_server_interface.proto",
}

// AgentHandlerClient is the client API for AgentHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentHandlerClient interface {
	RegisterAgent(ctx context.Context, in *ClientDetails, opts ...grpc.CallOption) (*ClientCredentials, error)
}

type agentHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentHandlerClient(cc grpc.ClientConnInterface) AgentHandlerClient {
	return &agentHandlerClient{cc}
}

func (c *agentHandlerClient) RegisterAgent(ctx context.Context, in *ClientDetails, opts ...grpc.CallOption) (*ClientCredentials, error) {
	out := new(ClientCredentials)
	err := c.cc.Invoke(ctx, "/experiment.AgentHandler/RegisterAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentHandlerServer is the server API for AgentHandler service.
// All implementations must embed UnimplementedAgentHandlerServer
// for forward compatibility
type AgentHandlerServer interface {
	RegisterAgent(context.Context, *ClientDetails) (*ClientCredentials, error)
	mustEmbedUnimplementedAgentHandlerServer()
}

// UnimplementedAgentHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAgentHandlerServer struct {
}

func (UnimplementedAgentHandlerServer) RegisterAgent(context.Context, *ClientDetails) (*ClientCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedAgentHandlerServer) mustEmbedUnimplementedAgentHandlerServer() {}

// UnsafeAgentHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentHandlerServer will
// result in compilation errors.
type UnsafeAgentHandlerServer interface {
	mustEmbedUnimplementedAgentHandlerServer()
}

func RegisterAgentHandlerServer(s grpc.ServiceRegistrar, srv AgentHandlerServer) {
	s.RegisterService(&AgentHandler_ServiceDesc, srv)
}

func _AgentHandler_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentHandlerServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.AgentHandler/RegisterAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentHandlerServer).RegisterAgent(ctx, req.(*ClientDetails))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentHandler_ServiceDesc is the grpc.ServiceDesc for AgentHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiment.AgentHandler",
	HandlerType: (*AgentHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAgent",
			Handler:    _AgentHandler_RegisterAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ml_server_interface.proto",
}

// ExperimentHandlerClient is the client API for ExperimentHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperimentHandlerClient interface {
	Setup(ctx context.Context, in *ExperimentSetupRequest, opts ...grpc.CallOption) (*ExperimentSetupResponse, error)
	StartExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*Null, error)
	StatusExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*ExperimentSetupResponse, error)
	StopExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*TerminationResp, error)
}

type experimentHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentHandlerClient(cc grpc.ClientConnInterface) ExperimentHandlerClient {
	return &experimentHandlerClient{cc}
}

func (c *experimentHandlerClient) Setup(ctx context.Context, in *ExperimentSetupRequest, opts ...grpc.CallOption) (*ExperimentSetupResponse, error) {
	out := new(ExperimentSetupResponse)
	err := c.cc.Invoke(ctx, "/experiment.ExperimentHandler/Setup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentHandlerClient) StartExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/experiment.ExperimentHandler/StartExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentHandlerClient) StatusExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*ExperimentSetupResponse, error) {
	out := new(ExperimentSetupResponse)
	err := c.cc.Invoke(ctx, "/experiment.ExperimentHandler/StatusExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentHandlerClient) StopExperiment(ctx context.Context, in *Experiment, opts ...grpc.CallOption) (*TerminationResp, error) {
	out := new(TerminationResp)
	err := c.cc.Invoke(ctx, "/experiment.ExperimentHandler/StopExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentHandlerServer is the server API for ExperimentHandler service.
// All implementations must embed UnimplementedExperimentHandlerServer
// for forward compatibility
type ExperimentHandlerServer interface {
	Setup(context.Context, *ExperimentSetupRequest) (*ExperimentSetupResponse, error)
	StartExperiment(context.Context, *Experiment) (*Null, error)
	StatusExperiment(context.Context, *Experiment) (*ExperimentSetupResponse, error)
	StopExperiment(context.Context, *Experiment) (*TerminationResp, error)
	mustEmbedUnimplementedExperimentHandlerServer()
}

// UnimplementedExperimentHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentHandlerServer struct {
}

func (UnimplementedExperimentHandlerServer) Setup(context.Context, *ExperimentSetupRequest) (*ExperimentSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (UnimplementedExperimentHandlerServer) StartExperiment(context.Context, *Experiment) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartExperiment not implemented")
}
func (UnimplementedExperimentHandlerServer) StatusExperiment(context.Context, *Experiment) (*ExperimentSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusExperiment not implemented")
}
func (UnimplementedExperimentHandlerServer) StopExperiment(context.Context, *Experiment) (*TerminationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopExperiment not implemented")
}
func (UnimplementedExperimentHandlerServer) mustEmbedUnimplementedExperimentHandlerServer() {}

// UnsafeExperimentHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentHandlerServer will
// result in compilation errors.
type UnsafeExperimentHandlerServer interface {
	mustEmbedUnimplementedExperimentHandlerServer()
}

func RegisterExperimentHandlerServer(s grpc.ServiceRegistrar, srv ExperimentHandlerServer) {
	s.RegisterService(&ExperimentHandler_ServiceDesc, srv)
}

func _ExperimentHandler_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExperimentSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentHandlerServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.ExperimentHandler/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentHandlerServer).Setup(ctx, req.(*ExperimentSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentHandler_StartExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Experiment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentHandlerServer).StartExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.ExperimentHandler/StartExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentHandlerServer).StartExperiment(ctx, req.(*Experiment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentHandler_StatusExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Experiment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentHandlerServer).StatusExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.ExperimentHandler/StatusExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentHandlerServer).StatusExperiment(ctx, req.(*Experiment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentHandler_StopExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Experiment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentHandlerServer).StopExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.ExperimentHandler/StopExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentHandlerServer).StopExperiment(ctx, req.(*Experiment))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperimentHandler_ServiceDesc is the grpc.ServiceDesc for ExperimentHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperimentHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiment.ExperimentHandler",
	HandlerType: (*ExperimentHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _ExperimentHandler_Setup_Handler,
		},
		{
			MethodName: "StartExperiment",
			Handler:    _ExperimentHandler_StartExperiment_Handler,
		},
		{
			MethodName: "StatusExperiment",
			Handler:    _ExperimentHandler_StatusExperiment_Handler,
		},
		{
			MethodName: "StopExperiment",
			Handler:    _ExperimentHandler_StopExperiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ml_server_interface.proto",
}

// IterationCallBacksHandlerClient is the client API for IterationCallBacksHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IterationCallBacksHandlerClient interface {
	IterationsResults(ctx context.Context, in *IterationResp, opts ...grpc.CallOption) (*Null, error)
}

type iterationCallBacksHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewIterationCallBacksHandlerClient(cc grpc.ClientConnInterface) IterationCallBacksHandlerClient {
	return &iterationCallBacksHandlerClient{cc}
}

func (c *iterationCallBacksHandlerClient) IterationsResults(ctx context.Context, in *IterationResp, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/experiment.IterationCallBacksHandler/IterationsResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IterationCallBacksHandlerServer is the server API for IterationCallBacksHandler service.
// All implementations must embed UnimplementedIterationCallBacksHandlerServer
// for forward compatibility
type IterationCallBacksHandlerServer interface {
	IterationsResults(context.Context, *IterationResp) (*Null, error)
	mustEmbedUnimplementedIterationCallBacksHandlerServer()
}

// UnimplementedIterationCallBacksHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedIterationCallBacksHandlerServer struct {
}

func (UnimplementedIterationCallBacksHandlerServer) IterationsResults(context.Context, *IterationResp) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IterationsResults not implemented")
}
func (UnimplementedIterationCallBacksHandlerServer) mustEmbedUnimplementedIterationCallBacksHandlerServer() {
}

// UnsafeIterationCallBacksHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IterationCallBacksHandlerServer will
// result in compilation errors.
type UnsafeIterationCallBacksHandlerServer interface {
	mustEmbedUnimplementedIterationCallBacksHandlerServer()
}

func RegisterIterationCallBacksHandlerServer(s grpc.ServiceRegistrar, srv IterationCallBacksHandlerServer) {
	s.RegisterService(&IterationCallBacksHandler_ServiceDesc, srv)
}

func _IterationCallBacksHandler_IterationsResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IterationResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IterationCallBacksHandlerServer).IterationsResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/experiment.IterationCallBacksHandler/IterationsResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IterationCallBacksHandlerServer).IterationsResults(ctx, req.(*IterationResp))
	}
	return interceptor(ctx, in, info, handler)
}

// IterationCallBacksHandler_ServiceDesc is the grpc.ServiceDesc for IterationCallBacksHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IterationCallBacksHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiment.IterationCallBacksHandler",
	HandlerType: (*IterationCallBacksHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IterationsResults",
			Handler:    _IterationCallBacksHandler_IterationsResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ml_server_interface.proto",
}
