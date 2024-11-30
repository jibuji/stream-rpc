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

func (s *CalculatorService) Add(ctx context.Context, req *proto.AddRequest) *proto.AddResponse {
	// TODO: Implement your logic here
	return &proto.AddResponse{Result: req.A + req.B}
}
func (s *CalculatorService) Multiply(ctx context.Context, req *proto.MultiplyRequest) *proto.MultiplyResponse {
	// TODO: Implement your logic here
	return &proto.MultiplyResponse{Result: req.A * req.B}
}
func (s *CalculatorService) Divide(ctx context.Context, req *proto.DivideRequest) *proto.DivideResponse {
	// TODO: Implement your logic here
	return &proto.DivideResponse{Result: req.A / req.B}
}
