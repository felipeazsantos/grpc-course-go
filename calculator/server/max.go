package main

import (
	"io"
	"log"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {

	max := int64(0)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading Max stream: %v\n", err)
		}

		if max < int64(msg.Number) {
			max = int64(msg.Number)
		}

		err = stream.Send(&pb.MaxResponse{
			Result: max,
		})

		if err != nil {
			log.Fatalf("Error while sending Max to the client: %v\n", err)
		}

	}

	return nil
}
