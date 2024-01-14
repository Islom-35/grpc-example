package server

import (
	"fmt"
	"log"
	"net"

	pb "imantask/internal/genproto/pb"

	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv    *grpc.Server
	postServer pb.CollectorServiceServer
}

func New(postServer pb.CollectorServiceServer) Server {
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

	pb.RegisterCollectorServiceServer(s.grpcSrv, s.postServer)
	log.Println("listening on port: ", addr)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
