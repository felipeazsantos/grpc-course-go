package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/felipeazsantos/grpc-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {

	numbers := []int64{1, 5, 3, 6, 2, 20}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			err = stream.Send(&pb.MaxRequest{
				Number: number,
			})
			if err != nil {
				log.Fatalf("Error sending Max request: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving Max server message: %v\n", err)
				break
			}

			log.Printf("Max numbers: %d", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
