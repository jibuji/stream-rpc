// Code generated by protoc-gen-stream-rpc. DO NOT EDIT.
package service

import (
	"context"
	proto "github.com/jibuji/go-stream-rpc/examples/calculator/proto"
)

// CalculatorService implements the Calculator service
type CalculatorService struct {
	proto.UnimplementedCalculatorServer
}

// Add implements CalculatorServer
func (s *CalculatorService) Add(ctx context.Context, req *proto.AddRequest) *proto.AddResponse {
	// TODO: Implement your logic here
	return &proto.AddResponse{}
}

// Multiply implements CalculatorServer
func (s *CalculatorService) Multiply(ctx context.Context, req *proto.MultiplyRequest) *proto.MultiplyResponse {
	// TODO: Implement your logic here
	return &proto.MultiplyResponse{}
}
