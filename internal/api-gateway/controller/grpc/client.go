package grpc

import (
	"context"
	"fmt"
	"log"

	pb "imantask/internal/genproto/pb"
	ppb"imantask/internal/genproto/ppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	collectClient pb.CollectorServiceClient
	postClient    ppb.PostServiceClient
}

func NewClient(port, port2 int) (*Client, error) {
	var conn *grpc.ClientConn
	var conn2 *grpc.ClientConn

	addr := fmt.Sprintf("iman_task_collector_app_1:%d", port)
	addr2:=fmt.Sprintf("iman_task_post_app_1:%d", port2)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	conn2, err =grpc.Dial(addr2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:          conn,
		collectClient: pb.NewCollectorServiceClient(conn),
		postClient: ppb.NewPostServiceClient(conn2),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) CollectPosts(ctx context.Context) error {
	_, err := c.collectClient.CollectPosts(ctx, &pb.CollectPostsRequest{})
	if err != nil {
		log.Printf("Collect failed: %v", err)
		return err
	}

	return nil
}

func (c *Client) GetPostByID(ctx context.Context,ID *ppb.ID)(*ppb.PostResponse,error){
	post,err:= c.postClient.GetPostByID(ctx,ID)
	if err!=nil{
		log.Printf("GetByID(): %v",err)
		return &ppb.PostResponse{},err
	}
	return post,nil
}

func (c *Client) GetPage(ctx context.Context,Page *ppb.Page)(*ppb.PostResponseList,error){
	posts, err:=c.postClient.GetPage(ctx,Page)
	if err!=nil{
		log.Printf("GetByID(): %v",err)
		return &ppb.PostResponseList{},err
	}
	return posts,nil
}

func (c *Client) DeletePostByID(ctx context.Context,ID *ppb.ID)error{
	_, err:= c.postClient.DeleteByID(ctx,ID)
	if err!=nil{
		log.Printf("GetByID(): %v",err)
		return err
	}
	return nil
}
