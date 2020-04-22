package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/usecase/ucmock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"github.com/stretchr/testify/assert"
	"testing"
)

type orderTest struct {
	ctrl *gomock.Controller
	uc   *ucmock.MockOrderUC
	unit *OrderService
	ctx  context.Context
}

func newOrderTest() *orderTest {
	return &orderTest{}
}

func (ot *orderTest) Begin(t *testing.T) {
	ot.ctrl = gomock.NewController(t)
	ot.uc = ucmock.NewMockOrderUC(ot.ctrl)
	ot.unit = NewOrderService(ot.uc)
	ot.ctx = context.Background()
}

func (ot *orderTest) Finish() {
	ot.ctrl.Finish()
}

func TestOrder_List(t *testing.T) {
	test := newOrderTest()
	t.Run("given nil param, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.ListOrder(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.ListOrderReq{
			CustomerId: 1,
			ShopId:     2,
		}
		test.uc.EXPECT().List(request.ListOrderReq{UserID: req.GetCustomerId(), ShopID: req.GetShopId()}).Return(nil, errors.New("error"))

		got, err := test.unit.ListOrder(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.ListOrderReq{
			CustomerId: 1,
			ShopId:     2,
		}
		test.uc.EXPECT().List(request.ListOrderReq{UserID: req.GetCustomerId(), ShopID: req.GetShopId()}).Return([]entity.Order{
			{
				ID:     1,
				UserID: 1,
				ShopID: 2,
			},
			{
				ID:     2,
				UserID: 1,
				ShopID: 2,
			},
		}, nil)

		got, err := test.unit.ListOrder(context.Background(), req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}

func TestOrderService_GetOrder(t *testing.T) {
	test := newOrderTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.GetOrder(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.GetOrderReq{
			OrderId:    1,
			CustomerId: 2,
			ShopId:     3,
		}
		test.uc.EXPECT().Get(entity.Order{
			ID:     1,
			UserID: 2,
			ShopID: 3,
		}).Return(entity.Order{}, errors.New("error"))

		got, err := test.unit.GetOrder(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.GetOrderReq{
			OrderId:    1,
			CustomerId: 2,
			ShopId:     3,
		}
		test.uc.EXPECT().Get(entity.Order{
			ID:     1,
			UserID: 2,
			ShopID: 3,
		}).Return(entity.Order{ID: 1}, nil)

		got, err := test.unit.GetOrder(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestOrderService_Checkout(t *testing.T) {
	test := newOrderTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.Checkout(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.CheckoutReq{
			UserId:        123,
			CartIds:       []int64{1, 2, 3},
			PaymentAmount: 123456,
		}

		test.uc.EXPECT().CreateFromCarts(request.CheckoutReq{
			UserID:        req.GetUserId(),
			CartIDs:       req.GetCartIds(),
			PaymentAmount: req.GetPaymentAmount(),
		}).Return(nil, errors.New("error"))

		got, err := test.unit.Checkout(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.CheckoutReq{
			UserId:        123,
			CartIds:       []int64{1, 2, 3},
			PaymentAmount: 123456,
		}

		test.uc.EXPECT().CreateFromCarts(request.CheckoutReq{
			CartIDs:       req.GetCartIds(),
			PaymentAmount: req.GetPaymentAmount(),
			UserID:        req.GetUserId(),
		}).Return([]entity.Order{{ID: 12337}}, nil)

		got, err := test.unit.Checkout(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}

func TestOrderService_Fulfill(t *testing.T) {
	test := newOrderTest()
	t.Run("given nil params should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.Fulfill(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.FulfillReq{
			OrderId: 123,
			UserId:  1412,
		}
		test.uc.EXPECT().Fulfill(entity.Order{ID: req.GetOrderId(), UserID: req.GetUserId()}).Return(entity.Order{}, errors.New("error"))
		got, err := test.unit.Fulfill(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.FulfillReq{
			OrderId: 123,
			UserId:  1412,
		}
		test.uc.EXPECT().Fulfill(entity.Order{ID: req.GetOrderId(), UserID: req.GetUserId()}).Return(entity.Order{ID: req.GetOrderId(), UserID: req.GetUserId()}, nil)
		got, err := test.unit.Fulfill(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
