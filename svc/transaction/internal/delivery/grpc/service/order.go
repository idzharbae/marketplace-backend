package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type OrderService struct {
	orderUC internal.OrderUC
}

func NewOrderService(orderUC internal.OrderUC) *OrderService {
	return &OrderService{orderUC: orderUC}
}

func (os *OrderService) Checkout(ctx context.Context, in *prototransaction.CheckoutReq) (*prototransaction.CheckoutResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	order, err := os.orderUC.CreateFromCarts(request.CheckoutReq{
		CartIDs:       in.GetCartIds(),
		PaymentAmount: in.GetPaymentAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.CheckoutResp{
		Order: converter.OrderEntityToProto(order),
	}, nil
}

func (os *OrderService) Fulfill(ctx context.Context, in *prototransaction.FulfillReq) (*prototransaction.FulfillResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	_, err := os.orderUC.Fulfill(entity.Order{
		ID:      in.GetOrderId(),
		Payment: entity.Payment{Amount: in.GetPaymentAmount()},
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.FulfillResp{
		Success: true,
	}, nil
}
