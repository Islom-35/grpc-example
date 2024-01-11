package server

import (
	"context"
	"imantask/internal/collector/app"
	"imantask/internal/genproto/pb"
)

type CollectorServer struct {
	service app.CollectorService
	pb.UnimplementedSaverServiceServer
}

func NewCollectorServer(service app.CollectorService) CollectorServer {
	return CollectorServer{
		service: service,
	}
}

func (h *CollectorServer) CollectPosts(ctx context.Context, req *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error) {
	return h.service.Save(ctx, req)
}
