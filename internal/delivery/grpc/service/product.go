package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/internal/app"
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
)

type ProductService struct {
}

func NewProductService(a *app.Marketplace) *ProductService {
	return &ProductService{}
}

func (p *ProductService) ListProducts(ctx context.Context, req *marketplaceproto.ListProductsReq) (*marketplaceproto.ListProductsResp, error) {
	return nil, nil
}
