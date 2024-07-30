// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: grpc/testing/benchmark_service.proto

package grpc_testing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BenchmarkService_UnaryCall_FullMethodName           = "/grpc.testing.BenchmarkService/UnaryCall"
	BenchmarkService_StreamingCall_FullMethodName       = "/grpc.testing.BenchmarkService/StreamingCall"
	BenchmarkService_StreamingFromClient_FullMethodName = "/grpc.testing.BenchmarkService/StreamingFromClient"
	BenchmarkService_StreamingFromServer_FullMethodName = "/grpc.testing.BenchmarkService/StreamingFromServer"
	BenchmarkService_StreamingBothWays_FullMethodName   = "/grpc.testing.BenchmarkService/StreamingBothWays"
)

// BenchmarkServiceClient is the client API for BenchmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BenchmarkServiceClient interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// Repeated sequence of one request followed by one response.
	// Should be called streaming ping-pong
	// The server returns the client payload as-is on each response
	StreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error)
	// Single-sided unbounded streaming from client to server
	// The server returns the client payload as-is once the client does WritesDone
	StreamingFromClient(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SimpleRequest, SimpleResponse], error)
	// Single-sided unbounded streaming from server to client
	// The server repeatedly returns the client payload as-is
	StreamingFromServer(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SimpleResponse], error)
	// Two-sided unbounded streaming between server to client
	// Both sides send the content of their own choice to the other
	StreamingBothWays(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error)
}

type benchmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkServiceClient(cc grpc.ClientConnInterface) BenchmarkServiceClient {
	return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, BenchmarkService_UnaryCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchmarkServiceClient) StreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BenchmarkService_ServiceDesc.Streams[0], BenchmarkService_StreamingCall_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SimpleRequest, SimpleResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingCallClient = grpc.BidiStreamingClient[SimpleRequest, SimpleResponse]

func (c *benchmarkServiceClient) StreamingFromClient(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SimpleRequest, SimpleResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BenchmarkService_ServiceDesc.Streams[1], BenchmarkService_StreamingFromClient_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SimpleRequest, SimpleResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingFromClientClient = grpc.ClientStreamingClient[SimpleRequest, SimpleResponse]

func (c *benchmarkServiceClient) StreamingFromServer(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SimpleResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BenchmarkService_ServiceDesc.Streams[2], BenchmarkService_StreamingFromServer_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SimpleRequest, SimpleResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingFromServerClient = grpc.ServerStreamingClient[SimpleResponse]

func (c *benchmarkServiceClient) StreamingBothWays(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BenchmarkService_ServiceDesc.Streams[3], BenchmarkService_StreamingBothWays_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SimpleRequest, SimpleResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingBothWaysClient = grpc.BidiStreamingClient[SimpleRequest, SimpleResponse]

// BenchmarkServiceServer is the server API for BenchmarkService service.
// All implementations must embed UnimplementedBenchmarkServiceServer
// for forward compatibility.
type BenchmarkServiceServer interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// Repeated sequence of one request followed by one response.
	// Should be called streaming ping-pong
	// The server returns the client payload as-is on each response
	StreamingCall(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error
	// Single-sided unbounded streaming from client to server
	// The server returns the client payload as-is once the client does WritesDone
	StreamingFromClient(grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]) error
	// Single-sided unbounded streaming from server to client
	// The server repeatedly returns the client payload as-is
	StreamingFromServer(*SimpleRequest, grpc.ServerStreamingServer[SimpleResponse]) error
	// Two-sided unbounded streaming between server to client
	// Both sides send the content of their own choice to the other
	StreamingBothWays(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error
	mustEmbedUnimplementedBenchmarkServiceServer()
}

// UnimplementedBenchmarkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBenchmarkServiceServer struct{}

func (UnimplementedBenchmarkServiceServer) UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (UnimplementedBenchmarkServiceServer) StreamingCall(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingCall not implemented")
}
func (UnimplementedBenchmarkServiceServer) StreamingFromClient(grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingFromClient not implemented")
}
func (UnimplementedBenchmarkServiceServer) StreamingFromServer(*SimpleRequest, grpc.ServerStreamingServer[SimpleResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingFromServer not implemented")
}
func (UnimplementedBenchmarkServiceServer) StreamingBothWays(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingBothWays not implemented")
}
func (UnimplementedBenchmarkServiceServer) mustEmbedUnimplementedBenchmarkServiceServer() {}
func (UnimplementedBenchmarkServiceServer) testEmbeddedByValue()                          {}

// UnsafeBenchmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BenchmarkServiceServer will
// result in compilation errors.
type UnsafeBenchmarkServiceServer interface {
	mustEmbedUnimplementedBenchmarkServiceServer()
}

func RegisterBenchmarkServiceServer(s grpc.ServiceRegistrar, srv BenchmarkServiceServer) {
	// If the following call pancis, it indicates UnimplementedBenchmarkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BenchmarkService_ServiceDesc, srv)
}

func _BenchmarkService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BenchmarkService_UnaryCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchmarkService_StreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).StreamingCall(&grpc.GenericServerStream[SimpleRequest, SimpleResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingCallServer = grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]

func _BenchmarkService_StreamingFromClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).StreamingFromClient(&grpc.GenericServerStream[SimpleRequest, SimpleResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingFromClientServer = grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]

func _BenchmarkService_StreamingFromServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BenchmarkServiceServer).StreamingFromServer(m, &grpc.GenericServerStream[SimpleRequest, SimpleResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingFromServerServer = grpc.ServerStreamingServer[SimpleResponse]

func _BenchmarkService_StreamingBothWays_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).StreamingBothWays(&grpc.GenericServerStream[SimpleRequest, SimpleResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BenchmarkService_StreamingBothWaysServer = grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]

// BenchmarkService_ServiceDesc is the grpc.ServiceDesc for BenchmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BenchmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _BenchmarkService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingCall",
			Handler:       _BenchmarkService_StreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingFromClient",
			Handler:       _BenchmarkService_StreamingFromClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingFromServer",
			Handler:       _BenchmarkService_StreamingFromServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingBothWays",
			Handler:       _BenchmarkService_StreamingBothWays_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/testing/benchmark_service.proto",
}
