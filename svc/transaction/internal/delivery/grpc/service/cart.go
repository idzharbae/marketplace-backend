package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type CartService struct {
	cartUC internal.CartUC
}

func NewCartService(cartUC internal.CartUC) *CartService {
	return &CartService{cartUC: cartUC}
}

func (cs *CartService) ListCartItems(ctx context.Context, in *prototransaction.ListCartItemsReq) (*prototransaction.ListCartItemsResp, error) {
	if in == nil {
		return nil, errors.New("param should not be nil")
	}
	res, err := cs.cartUC.List(in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &prototransaction.ListCartItemsResp{
		Cart: converter.CartEntitiesToProtos(res),
	}, nil
}

func (cs *CartService) AddToCart(ctx context.Context, in *prototransaction.AddToCartReq) (*prototransaction.Cart, error) {
	if in == nil {
		return nil, errors.New("param should not be nil")
	}
	res, err := cs.cartUC.Add(entity.Cart{
		Product: entity.Product{
			ID: in.GetProductId(),
		},
		UserID:   in.GetUserId(),
		AmountKG: in.GetQuantityKg(),
	})
	if err != nil {
		return nil, err
	}
	return converter.CartEntityToProto(res), nil
}

func (cs *CartService) UpdateCart(ctx context.Context, in *prototransaction.UpdateCartReq) (*prototransaction.Cart, error) {
	if in == nil {
		return nil, errors.New("param should not be nil")
	}
	res, err := cs.cartUC.Update(entity.Cart{
		ID:       in.GetId(),
		AmountKG: in.GetQuantityKg(),
		UserID:   in.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return converter.CartEntityToProto(res), nil
}

func (cs *CartService) RemoveCart(ctx context.Context, in *prototransaction.RemoveCartReq) (*prototransaction.RemoveCartResp, error) {
	if in == nil {
		return nil, errors.New("param should not be nil")
	}
	err := cs.cartUC.Remove(in.GetId(), in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &prototransaction.RemoveCartResp{Success: true}, nil
}
