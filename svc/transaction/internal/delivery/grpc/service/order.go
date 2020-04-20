package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type OrderService struct {
	orderUC internal.OrderUC
	cartUC  internal.CartUC
}

func NewOrderService(orderUC internal.OrderUC) *OrderService {
	return &OrderService{orderUC: orderUC}
}

func (os *OrderService) Checkout(ctx context.Context, in *prototransaction.CheckoutReq) (*prototransaction.CheckoutResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	return nil, nil
}

func (os *OrderService) Fulfill(ctx context.Context, in *prototransaction.FulfillReq) (*prototransaction.FulfillResp, error) {
	return nil, nil
}
