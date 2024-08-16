package main

import (
	"context"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
)

func (server *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
