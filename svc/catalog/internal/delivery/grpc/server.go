package grpc

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	config     *Config
}

// Config defines config for GRPC service
type Config struct {
	ServerOptions      []grpc.ServerOption
	Port               string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

// New creates new grpc service
func NewServer(port string, opts ...Option) *Server {
	cfg := Config{Port: port}

	// default unary interceptors
	cfg.unaryInterceptors = append(cfg.unaryInterceptors,
		grpc_prometheus.UnaryServerInterceptor,
		grpc_validator.UnaryServerInterceptor(),
	)

	// default stream interceptors
	cfg.streamInterceptors = append(cfg.streamInterceptors,
		grpc_prometheus.StreamServerInterceptor,
	)

	for _, opt := range opts {
		cfg = opt(cfg)
	}

	cfg.ServerOptions = append(cfg.ServerOptions,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(cfg.streamInterceptors...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(cfg.unaryInterceptors...)),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()
	s := grpc.NewServer(cfg.ServerOptions...)
	grpc_prometheus.Register(s)
	reflection.Register(s)

	return &Server{
		grpcServer: s,
		config:     &cfg,
	}
}

// Config return grpc service config
func (svc *Server) Config() *Config {
	return svc.config
}

// Address of grpc service
func (svc *Server) Port() string {
	return svc.config.Port
}

// Type of grpc service
func (svc *Server) Type() string {
	svcname := "grpc-service"
	return svcname
}

// Serve grpc server
func (svc *Server) Serve(ls net.Listener) error {
	return svc.Server().Serve(ls)
}

// Shutdown grpc server
func (svc *Server) Shutdown(ctx context.Context) error {
	svc.Server().GracefulStop()
	return nil
}

// Server return grpc server
func (svc *Server) Server() *grpc.Server {
	return svc.grpcServer
}

// Code returns service code to be used by admin page
func (svc Server) Code() string {
	return "grpc"
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		return err
	}
	log.Printf("listening to port " + s.Port())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
