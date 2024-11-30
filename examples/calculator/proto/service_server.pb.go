// Code generated by stream-rpc. DO NOT EDIT.
package proto

import (
	rpc "github.com/jibuji/go-stream-rpc"
	"context"
)

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations
type UnimplementedCalculatorServer struct{}

type CalculatorServer interface {
	Add(context.Context, *AddRequest) *AddResponse

	Multiply(context.Context, *MultiplyRequest) *MultiplyResponse

	Divide(context.Context, *DivideRequest) *DivideResponse
}

type CalculatorServerImpl struct {
	impl CalculatorServer
}

func RegisterCalculatorServer(peer *rpc.RpcPeer, impl CalculatorServer) {
	server := &CalculatorServerImpl{impl: impl}
	peer.RegisterService("Calculator", server)
}

func (s *UnimplementedCalculatorServer) Add(ctx context.Context, req *AddRequest) *AddResponse {
	return nil
}

func (s *UnimplementedCalculatorServer) Multiply(ctx context.Context, req *MultiplyRequest) *MultiplyResponse {
	return nil
}

func (s *UnimplementedCalculatorServer) Divide(ctx context.Context, req *DivideRequest) *DivideResponse {
	return nil
}

func (s *CalculatorServerImpl) Add(ctx context.Context, req *AddRequest) *AddResponse {
	return s.impl.Add(ctx, req)
}

func (s *CalculatorServerImpl) Multiply(ctx context.Context, req *MultiplyRequest) *MultiplyResponse {
	return s.impl.Multiply(ctx, req)
}

func (s *CalculatorServerImpl) Divide(ctx context.Context, req *DivideRequest) *DivideResponse {
	return s.impl.Divide(ctx, req)
}
