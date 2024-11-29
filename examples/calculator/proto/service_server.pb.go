// Code generated by stream-rpc. DO NOT EDIT.
package proto

import (
	rpc "stream-rpc"
)

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations
type UnimplementedCalculatorServer struct{}

func (s *UnimplementedCalculatorServer) Add(*AddRequest) (*AddResponse, error) {
	return nil, rpc.ErrNotImplemented
}

func (s *UnimplementedCalculatorServer) Multiply(*MultiplyRequest) (*MultiplyResponse, error) {
	return nil, rpc.ErrNotImplemented
}

type CalculatorServer interface {
	Add(*AddRequest) (*AddResponse, error)

	Multiply(*MultiplyRequest) (*MultiplyResponse, error)
}

type CalculatorServerImpl struct {
	impl CalculatorServer
}

func RegisterCalculatorServer(peer *rpc.RpcPeer, impl CalculatorServer) {
	server := &CalculatorServerImpl{impl: impl}
	peer.RegisterService("Calculator", server)
}

func (s *CalculatorServerImpl) Add(req *AddRequest) (*AddResponse, error) {
	return s.impl.Add(req)
}

func (s *CalculatorServerImpl) Multiply(req *MultiplyRequest) (*MultiplyResponse, error) {
	return s.impl.Multiply(req)
}