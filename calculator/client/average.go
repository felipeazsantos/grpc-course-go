package main

import (
	"context"
	"log"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {

	reqs := []*pb.AverageRequest{
		{Num: 10},
		{Num: 5},
		{Num: 8},
		{Num: 6},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error calling the average method: %v", err)
	}

	for _, req := range reqs {

		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Error sending average request: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving average response: %v\n", err)
	}

	log.Printf("The Average is: %.2f", res.Result)
}
