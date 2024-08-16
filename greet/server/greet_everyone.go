package main

import (
	"io"
	"log"

	pb "github.com/felipeazsantos/grpc-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while seding data to client: %v\n", err)
		}
	}
}
