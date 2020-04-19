package connection

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"google.golang.org/grpc"
)

type Catalog interface {
	GetProduct(ctx context.Context, in *catalogproto.GetProductReq, opts ...grpc.CallOption) (*catalogproto.Product, error)
	Close() error
}
