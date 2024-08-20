package main

import (
	"context"
	"io"
	"log"

	pb "github.com/felipeazsantos/grpc-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("listBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listBlogs stream: %v\n", err)
	}

	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading listBlogs: %v\n", err)
		}

		log.Printf("Blog: %v", blog)
	}
}
