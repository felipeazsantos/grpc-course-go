package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doCalculator(c)
	// doPrimes(c)
	// doAverage(c)
	// doMax(c)
	// doSqrt(c, 10)
	doSqrt(c, -10)
}
