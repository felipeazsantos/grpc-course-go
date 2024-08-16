package main

import (
	"context"
	"log"

	pb "github.com/felipeazsantos/grpc-course/greet/proto"
)

func (server *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
