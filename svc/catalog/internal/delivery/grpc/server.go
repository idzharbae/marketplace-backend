package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	port       string
}

func NewServer(port string) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		port:       port,
	}
}

func (svc *Server) Server() *grpc.Server {
	return svc.grpcServer
}
func (svc *Server) Serve(ls net.Listener) error {
	return svc.Server().Serve(ls)
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	log.Printf("listening to port " + s.port)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
