package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type ProductService struct {
	ProductUC internal.ProductUC
}

func NewProductService(productUC internal.ProductUC) *ProductService {
	return &ProductService{
		ProductUC: productUC,
	}
}

func (p *ProductService) ListProducts(ctx context.Context, req *catalogproto.ListProductsReq) (*catalogproto.ListProductsResp, error) {
	ucReq := requests.ListProduct{Pagination: requests.Pagination{
		Page:  int(req.Pagination.Page),
		Limit: int(req.Pagination.Limit),
	}}
	products, err := p.ProductUC.List(ucReq)
	productProtos := converter.ProductEntitiesToProtos(products)

	if err != nil {
		return nil, err
	}

	return &catalogproto.ListProductsResp{
		Products: productProtos,
	}, nil
}

func (p *ProductService) GetProductByID(ctx context.Context, req *catalogproto.GetProductByIDReq) (*catalogproto.Product, error) {
	got, err := p.ProductUC.GetByID(req.GetId())
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}

func (p *ProductService) GetProductBySlug(ctx context.Context, req *catalogproto.GetProductBySlugReq) (*catalogproto.Product, error) {
	got, err := p.ProductUC.GetBySlug(req.Slug)
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}

func (p *ProductService) CreateProduct(ctx context.Context, req *catalogproto.Product) (*catalogproto.Product, error) {
	product := converter.ProductProtoToEntity(req)
	got, err := p.ProductUC.Create(product)
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *catalogproto.Product) (*catalogproto.Product, error) {
	product := converter.ProductProtoToEntity(req)
	got, err := p.ProductUC.Update(product)
	if err != nil {
		return nil, err
	}
	return converter.ProductEntityToProto(got), nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, req *catalogproto.PKReq) (*catalogproto.Empty, error) {
	return &catalogproto.Empty{}, p.ProductUC.Delete(req.GetId())
}
