package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
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
	ucReq := requests.ListProduct{
		ShopIDs:    req.GetShopIDs(),
		Category:   req.GetCategory(),
		Search:     req.GetSearch(),
		OrderBy:    req.GetOrderBy(),
		OrderType:  req.GetOrderType(),
		ProductIDs: req.GetProductIds(),
		Pagination: requests.Pagination{
			Page:  int(req.GetPagination().GetPage()),
			Limit: int(req.GetPagination().GetLimit()),
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

func (p *ProductService) GetProduct(ctx context.Context, req *catalogproto.GetProductReq) (*catalogproto.Product, error) {
	productReq := entity.Product{
		ID:   req.GetId(),
		Slug: req.GetSlug(),
	}
	got, err := p.ProductUC.Get(productReq)
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

func (p *ProductService) DeleteProduct(ctx context.Context, req *catalogproto.GetProductReq) (*catalogproto.Empty, error) {
	product := entity.Product{
		ID:   req.GetId(),
		Slug: req.GetSlug(),
	}
	return &catalogproto.Empty{}, p.ProductUC.Delete(product)
}
