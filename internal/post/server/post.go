package server

import (
	"context"
	pb 	"imantask/internal/genproto/ppb"
	"imantask/internal/post/app"
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

func (h *PostServer)UpdatePostByID(ctx context.Context, req *pb.UpdateRequest)(*pb.Empty,error){
	return h.service.UpdatePostByID(ctx,req)
}

func (h *PostServer) DeleteByID(ctx context.Context, req *pb.ID)(*pb.Empty,error){
	return h.service.DeletePostByID(ctx,req)
}