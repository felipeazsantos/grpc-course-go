package main

import (
	"io"
	"log"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.AverageResponse]) error {
	log.Printf("Average function was invoked with: %v\n", stream)

	count := 0
	sum := 0.0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: sum / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("error while receiving average stream: %v\n", err)
		}

		sum += float64(res.Num)
		count++
	}
}
