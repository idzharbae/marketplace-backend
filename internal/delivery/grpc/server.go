package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	config     *Config
}

// Config defines config for GRPC service
type Config struct {
	ServerOptions      []grpc.ServerOption
	Address            string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

// New creates new grpc service
func NewServer(address string, opts ...Option) *Server {
	cfg := Config{Address: address}

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
func (svc *Server) Server() *grpc.Server {
	return svc.grpcServer
}
