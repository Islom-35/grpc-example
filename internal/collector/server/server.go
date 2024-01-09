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
	postServer pb.SaverServiceServer
	pb.UnimplementedSaverServiceServer
}

func New(postServer pb.SaverServiceServer) Server {
	return Server{
		grpcSrv:    grpc.NewServer(),
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	pb.RegisterSaverServiceServer(s.grpcSrv, s.postServer)
	log.Println("listening on port: ", addr)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
