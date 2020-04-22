package connection

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"google.golang.org/grpc"
)

//go:generate mockgen -destination=connectionmock/catalog_mock.go -package=connectionmock github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection Catalog
type Catalog interface {
	ListProducts(ctx context.Context, in *catalogproto.ListProductsReq, opts ...grpc.CallOption) (*catalogproto.ListProductsResp, error)
	GetProduct(ctx context.Context, in *catalogproto.GetProductReq, opts ...grpc.CallOption) (*catalogproto.Product, error)
	Close() error
}

//go:generate mockgen -destination=connectionmock/auth_mock.go -package=connectionmock github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection Auth
type Auth interface {
	GetUser(ctx context.Context, in *authproto.GetUserReq, opts ...grpc.CallOption) (*authproto.User, error)
	UpdateSaldo(ctx context.Context, in *authproto.TopUpReq, opts ...grpc.CallOption) (*authproto.TopUpResp, error)
	Close() error
}
