package main

import (
	"context"
	"log"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})
	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}
	log.Printf("The sum is: %d", res.Result)
}
