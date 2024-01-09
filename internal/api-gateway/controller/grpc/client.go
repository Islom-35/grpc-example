package grpc

import (
	"context"
	"fmt"
	"log"

	pb 	"imantask/internal/genproto/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn       *grpc.ClientConn
	postClient pb.SaverServiceClient
}

func NewClient(port int) (*Client, error) {
	var conn *grpc.ClientConn
	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:       conn,
		postClient: pb.NewSaverServiceClient(conn),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) Save(ctx context.Context) error {
	post :=pb.CollectPostsRequest{}
	response, err := c.postClient.CollectPosts(ctx, &post)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
		return err
	}
	

	log.Printf("Save response: %v", response)
	return nil
}
