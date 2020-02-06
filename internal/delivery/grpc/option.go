package grpc

import "google.golang.org/grpc"

type Option func(Config) Config

// WithUnaryInterceptors add unary interceptor(s) to the grpc server
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(cfg Config) Config {
		cfg.unaryInterceptors = append(cfg.unaryInterceptors, interceptors...)

		return cfg
	}
}
