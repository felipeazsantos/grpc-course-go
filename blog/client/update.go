package main

import (
	"context"
	"log"

	pb "github.com/felipeazsantos/grpc-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "Not Felipe",
		Title:    "A new title",
		Content:  "Content of first blog post, with awesome additions",
	}

	_, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while updating blog: %v\n", err)
		return
	}

	log.Println("Blog was updated")
}
