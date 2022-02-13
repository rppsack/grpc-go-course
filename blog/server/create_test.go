package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success", func(mt *mtest.T) {
		collection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		blog := &pb.Blog{
			AuthorId: "Clement",
			Title:    "My First Blog",
			Content:  "Content of the first blog",
		}

		_, err := c.CreateBlog(context.Background(), blog)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestCreateError(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Error", func(mt *mtest.T) {
		collection = mt.Coll
		mt.AddMockResponses(bson.D{{"error", 0}})

		blog := &pb.Blog{
			AuthorId: "Clement",
			Title:    "My First Blog",
			Content:  "Content of the first blog",
		}

		_, err := c.CreateBlog(context.Background(), blog)

		if err == nil {
			t.Error("Expected error")
		}

		respErr, ok := status.FromError(err)

		if !ok {
			t.Error("Expected error")
		}

		if respErr.Code() != codes.Internal {
			t.Errorf("Expected Internal, got %v", respErr.Code().String())
		}
	})
}