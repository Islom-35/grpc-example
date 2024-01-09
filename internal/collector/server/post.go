package server

import (
	"context"
	"imantask/internal/collector/app"
	"imantask/internal/genproto/pb"
)

type PostServer struct {
	service app.PostService
	pb.UnimplementedSaverServiceServer
}


func NewPostServer(service app.PostService) PostServer {
	return PostServer{
		service: service,
	}
}

func (h *PostServer) CollectPosts(ctx context.Context, req *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error) {
	return h.service.Save(ctx, req)
}
