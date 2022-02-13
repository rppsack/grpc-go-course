package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")
	res, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		fmt.Printf("Error happened while deleting: %v \n", err)
	}

	fmt.Printf("Blog was deleted: %v \n", res)
}