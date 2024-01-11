package server

import (
	"context"
	pb 	"imantask/internal/genproto/ppb"
	"imantask/internal/post-grud/app"
)

type PostServer struct{
	service app.PostService
	pb.UnimplementedPostServiceServer
}

func NewPostServer (service app.PostService)PostServer{
	return PostServer{
		service: service,
	}
}

func (h *PostServer)GetPostByID(ctx context.Context,req *pb.ID)(*pb.PostResponse,error){
	
	return h.service.GetPostByID(ctx,req)
}

func (h *PostServer)GetPage(ctx context.Context,req *pb.Page)(*pb.PostResponseList,error){
	return h.service.GetPage(ctx,req)
}