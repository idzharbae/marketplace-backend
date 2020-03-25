package service

import (
	"context"

	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShopService struct {
	shopUC internal.ShopUC
}

func NewShopService(shopUC internal.ShopUC) *ShopService {
	return &ShopService{
		shopUC: shopUC,
	}
}

func (s *ShopService) ListShops(ctx context.Context, in *catalogproto.ListShopsReq) (*catalogproto.ListShopsResp, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "parameter should not be nil")
	}
	req := converter.ListShopProtoToReq(in)
	res, err := s.shopUC.List(req)
	if err != nil {
		return nil, err
	}
	shops := converter.ShopEntitiesToProtos(res)
	return &catalogproto.ListShopsResp{
		Shops: shops,
	}, nil
}

func (s *ShopService) GetShop(ctx context.Context, in *catalogproto.GetShopReq) (*catalogproto.Shop, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "parameter should not be nil")
	}
	shopReq := entity.Shop{
		ID:   in.GetId(),
		Slug: in.GetSlug(),
	}
	res, err := s.shopUC.Get(shopReq)
	if err != nil {
		return nil, err
	}
	shop := converter.ShopEntityToProto(res)
	return shop, nil
}

func (s *ShopService) CreateShop(ctx context.Context, in *catalogproto.Shop) (*catalogproto.Shop, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "parameter should not be nil")
	}
	res, err := s.shopUC.Create(converter.ShopProtoToEntity(in))
	if err != nil {
		return nil, err
	}
	shop := converter.ShopEntityToProto(res)
	return shop, nil
}
func (s *ShopService) UpdateShop(ctx context.Context, in *catalogproto.Shop) (*catalogproto.Shop, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "parameter should not be nil")
	}
	req := converter.ShopProtoToEntity(in)
	res, err := s.shopUC.Update(req)
	if err != nil {
		return nil, err
	}
	shop := converter.ShopEntityToProto(res)
	return shop, nil
}
func (s *ShopService) DeleteShop(ctx context.Context, in *catalogproto.GetShopReq) (*catalogproto.Empty, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "parameter should not be nil")
	}
	shopReq := entity.Shop{
		ID:   in.GetId(),
		Slug: in.GetSlug(),
	}
	err := s.shopUC.Delete(shopReq)
	if err != nil {
		return nil, err
	}
	return &catalogproto.Empty{}, nil
}
