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

func (os *OrderService) ListOrder(ctx context.Context, in *prototransaction.ListOrderReq) (*prototransaction.ListOrderResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	orders, err := os.orderUC.List(request.ListOrderReq{
		UserID: in.GetCustomerId(),
		ShopID: in.GetShopId(),
		Status: in.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.ListOrderResp{
		Orders: converter.OrderEntitiesToProtos(orders),
	}, nil
}
func (os *OrderService) GetOrder(ctx context.Context, in *prototransaction.GetOrderReq) (*prototransaction.GetOrderResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	order, err := os.orderUC.Get(entity.Order{
		ID:     in.GetOrderId(),
		UserID: in.GetCustomerId(),
		ShopID: in.GetShopId(),
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.GetOrderResp{
		Order: converter.OrderEntityToProto(order),
	}, nil
}

func (os *OrderService) Checkout(ctx context.Context, in *prototransaction.CheckoutReq) (*prototransaction.CheckoutResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	order, err := os.orderUC.CreateFromCarts(request.CheckoutReq{
		UserID:        in.GetUserId(),
		CartIDs:       in.GetCartIds(),
		PaymentAmount: in.GetPaymentAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.CheckoutResp{
		Orders: converter.OrderEntitiesToProtos(order),
	}, nil
}

func (os *OrderService) UpdateOrderStatusToOnShipment(ctx context.Context, in *prototransaction.ChangeProductStatusReq) (*prototransaction.ShipProductResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	order, err := os.orderUC.UpdateOrderStatusToOnShipment(in.GetOrderId(), in.GetShopId())
	if err != nil {
		return nil, err
	}
	return &prototransaction.ShipProductResp{
		Order: converter.OrderEntityToProto(order),
	}, nil
}

func (os *OrderService) RejectOrder(ctx context.Context, in *prototransaction.ChangeProductStatusReq) (*prototransaction.Order, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	order, err := os.orderUC.RejectOrder(in.GetOrderId(), in.GetShopId())
	if err != nil {
		return nil, err
	}
	return converter.OrderEntityToProto(order), nil
}

func (os *OrderService) Fulfill(ctx context.Context, in *prototransaction.FulfillReq) (*prototransaction.FulfillResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	_, err := os.orderUC.Fulfill(entity.Order{
		ID:     in.GetOrderId(),
		UserID: in.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &prototransaction.FulfillResp{
		Success: true,
	}, nil
}
