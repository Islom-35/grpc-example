package server

import (
	"fmt"
	pb "imantask/internal/genproto/ppb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv    *grpc.Server
	postServer pb.PostServiceServer
}

func New(postServer pb.PostServiceServer) Server {
	return Server{
		grpcSrv:    grpc.NewServer(),
		postServer: postServer,
	}
}

func (s *Server) ListenAndServe(port string) error {
	addr := fmt.Sprintf("%v", port)

	lis, err := net.Listen("tcp", addr)	
	if err != nil {
		return err
	}

	pb.RegisterPostServiceServer(s.grpcSrv, s.postServer)
	log.Println("listening on port: ", addr)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
