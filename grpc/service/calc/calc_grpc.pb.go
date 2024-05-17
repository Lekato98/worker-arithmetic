// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: calc.proto

package calc

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

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcClient interface {
	Equation(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Equation(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/Calc/Equation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
// All implementations must embed UnimplementedCalcServer
// for forward compatibility
type CalcServer interface {
	Equation(context.Context, *Input) (*Output, error)
	mustEmbedUnimplementedCalcServer()
}

// UnimplementedCalcServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (UnimplementedCalcServer) Equation(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Equation not implemented")
}
func (UnimplementedCalcServer) mustEmbedUnimplementedCalcServer() {}

// UnsafeCalcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServer will
// result in compilation errors.
type UnsafeCalcServer interface {
	mustEmbedUnimplementedCalcServer()
}

func RegisterCalcServer(s grpc.ServiceRegistrar, srv CalcServer) {
	s.RegisterService(&Calc_ServiceDesc, srv)
}

func _Calc_Equation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Equation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calc/Equation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Equation(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Calc_ServiceDesc is the grpc.ServiceDesc for Calc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Equation",
			Handler:    _Calc_Equation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
