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
	ucReq := requests.ListProduct{Pagination: requests.Pagination{
		Page:  int(req.Pagination.Page),
		Limit: int(req.Pagination.Limit),
	}}
	products, err := p.ProductUC.List(ucReq)
	productProtos := converter.ProductEntitiesToProtos(products)

	if err != nil {
		return nil, err
	}

	return &marketplaceproto.ListProductsResp{
		Products: productProtos,
	}, nil
}

func (p *ProductService) GetProductByID(ctx context.Context, req *marketplaceproto.GetProductByIDReq) (*marketplaceproto.Product, error) {
	got, err := p.ProductUC.GetByID(req.ID)
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}

func (p *ProductService) GetProductBySlug(ctx context.Context, req *marketplaceproto.GetProductBySlugReq) (*marketplaceproto.Product, error) {
	got, err := p.ProductUC.GetBySlug(req.Slug)
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}
