package pkg

import (
	"context"
	"grpcEx/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedCalculatorServiceServer
}

func (s *Server) Calculate(ctx context.Context, req *proto.CalculateRequest) (*proto.CalculateResponse, error) {
	a := req.GetFirstNumber()
	b := req.GetSecondNumber()

	var result float64
	switch req.GetOperation() {
	case proto.Operation_ADD:
		result = a + b
	case proto.Operation_SUBTRACT:
		result = a - b
	case proto.Operation_MULTIPLY:
		result = a * b
	case proto.Operation_DIVIDE:
		if b == 0 {
			return nil, status.Error(codes.InvalidArgument, "div on 0")
		}
		result = a / b

	default:
		return nil, status.Errorf(codes.InvalidArgument, "неизвестная операция: %v", req.GetOperation())
	}

	return &proto.CalculateResponse{Result: result}, nil
}
