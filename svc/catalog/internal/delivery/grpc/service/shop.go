package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
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
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *ShopService) GetShopByID(ctx context.Context, in *catalogproto.GetShopByPKReq) (*catalogproto.Shop, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
func (s *ShopService) GetShopBySlug(ctx context.Context, in *catalogproto.GetShopBySlugReq) (*catalogproto.Shop, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *ShopService) CreateShop(ctx context.Context, in *catalogproto.Shop) (*catalogproto.Shop, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
func (s *ShopService) UpdateShop(ctx context.Context, in *catalogproto.Shop) (*catalogproto.Shop, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
func (s *ShopService) DeleteShop(ctx context.Context, in *catalogproto.Shop) (*catalogproto.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
