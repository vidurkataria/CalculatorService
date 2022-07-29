// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: calcipb/calci.proto

package calcipb

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error)
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvgClient, error)
	MaxNum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxNumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/CalculatorService.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/CalculatorService.CalculatorService/Prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/CalculatorService.CalculatorService/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAvgClient{stream}
	return x, nil
}

type CalculatorService_ComputeAvgClient interface {
	Send(*CompAvgRequest) error
	CloseAndRecv() (*CompAvgResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAvgClient) Send(m *CompAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvgClient) CloseAndRecv() (*CompAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CompAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) MaxNum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxNumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/CalculatorService.CalculatorService/MaxNum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxNumClient{stream}
	return x, nil
}

type CalculatorService_MaxNumClient interface {
	Send(*MaxNumRequest) error
	Recv() (*MaxNumResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaxNumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxNumClient) Send(m *MaxNumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxNumClient) Recv() (*MaxNumResponse, error) {
	m := new(MaxNumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Prime(*PrimeRequest, CalculatorService_PrimeServer) error
	ComputeAvg(CalculatorService_ComputeAvgServer) error
	MaxNum(CalculatorService_MaxNumServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Prime(*PrimeRequest, CalculatorService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}
func (UnimplementedCalculatorServiceServer) ComputeAvg(CalculatorService_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}
func (UnimplementedCalculatorServiceServer) MaxNum(CalculatorService_MaxNumServer) error {
	return status.Errorf(codes.Unimplemented, "method MaxNum not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculatorService.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Prime(m, &calculatorServicePrimeServer{stream})
}

type CalculatorService_PrimeServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAvg(&calculatorServiceComputeAvgServer{stream})
}

type CalculatorService_ComputeAvgServer interface {
	SendAndClose(*CompAvgResponse) error
	Recv() (*CompAvgRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAvgServer) SendAndClose(m *CompAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvgServer) Recv() (*CompAvgRequest, error) {
	m := new(CompAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_MaxNum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).MaxNum(&calculatorServiceMaxNumServer{stream})
}

type CalculatorService_MaxNumServer interface {
	Send(*MaxNumResponse) error
	Recv() (*MaxNumRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxNumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxNumServer) Send(m *MaxNumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxNumServer) Recv() (*MaxNumRequest, error) {
	m := new(MaxNumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CalculatorService.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _CalculatorService_Prime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvg",
			Handler:       _CalculatorService_ComputeAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "MaxNum",
			Handler:       _CalculatorService_MaxNum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calcipb/calci.proto",
}
