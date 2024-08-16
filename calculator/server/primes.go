package main

import (
	"log"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("Primes function was invoked with: %v\n", in)
	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
