package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/internal"
	"github.com/idzharbae/marketplace-backend/internal/converter"
	"github.com/idzharbae/marketplace-backend/internal/requests"
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
)

type ProductService struct {
	ProductUC internal.ProductUC
}

func NewProductService(productUC internal.ProductUC) *ProductService {
	return &ProductService{
		ProductUC: productUC,
	}
}

func (p *ProductService) ListProducts(ctx context.Context, req *marketplaceproto.ListProductsReq) (*marketplaceproto.ListProductsResp, error) {
	products, err := p.ProductUC.List(requests.ListProduct{})
	productProtos := converter.ProductEntitiesToProtos(products)

	return &marketplaceproto.ListProductsResp{
		Products: productProtos,
	}, err
}
