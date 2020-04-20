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
			OrderId:       123,
			PaymentAmount: 1412,
		}
		test.uc.EXPECT().Fulfill(entity.Order{ID: req.GetOrderId(), Payment: entity.Payment{
			Amount: req.GetPaymentAmount()},
		}).Return(entity.Order{}, errors.New("error"))
		got, err := test.unit.Fulfill(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &prototransaction.FulfillReq{
			OrderId:       123,
			PaymentAmount: 1412,
		}
		test.uc.EXPECT().Fulfill(entity.Order{ID: req.GetOrderId(), Payment: entity.Payment{
			Amount: req.GetPaymentAmount()},
		}).Return(entity.Order{ID: req.GetOrderId(), Payment: entity.Payment{
			ID:            12,
			Amount:        req.GetPaymentAmount(),
			PaymentMethod: 1,
			PaymentStatus: 1,
		}}, nil)
		got, err := test.unit.Fulfill(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
