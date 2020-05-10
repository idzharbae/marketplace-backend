package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
)

type cartTest struct {
	ctrl *gomock.Controller
	uc   *ucmock.MockCartUC
	unit *CartService
	ctx  context.Context
}

func newCartTest() *cartTest {
	return &cartTest{}
}

func TestCartService_ListCartItems(t *testing.T) {
	test := newCartTest()
	t.Run("given nil param, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.ListCartItems(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetListReq()

		test.uc.EXPECT().List(request.ListCartReq{
			UserID: req.GetUserId(),
			Pagination: request.Pagination{
				Page:  int(req.GetPagination().GetPage()),
				Limit: int(req.GetPagination().GetLimit()),
			},
		}).Return(nil, errors.New("error"))

		got, err := test.unit.ListCartItems(test.ctx, req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetListReq()

		test.uc.EXPECT().List(request.ListCartReq{
			UserID: req.GetUserId(),
			Pagination: request.Pagination{
				Page:  int(req.GetPagination().GetPage()),
				Limit: int(req.GetPagination().GetLimit()),
			},
		}).Return(test.GetCarts(), nil)

		got, err := test.unit.ListCartItems(test.ctx, req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func TestCartService_AddToCart(t *testing.T) {
	test := newCartTest()
	t.Run("given nil param, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.AddToCart(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Add(entity.Cart{
			Product: entity.Product{
				ID: req.GetProduct().GetId(),
			},
			UserID:   req.GetUserId(),
			AmountKG: req.GetProduct().GetAmountKg(),
		}).Return(entity.Cart{}, errors.New("error"))

		got, err := test.unit.AddToCart(test.ctx, &prototransaction.AddToCartReq{
			ProductId:  req.GetProduct().GetId(),
			UserId:     req.GetUserId(),
			QuantityKg: req.GetProduct().GetAmountKg(),
		})
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Add(entity.Cart{
			Product: entity.Product{
				ID: req.GetProduct().GetId(),
			},
			UserID:   req.GetUserId(),
			AmountKG: req.GetProduct().GetAmountKg(),
		}).Return(test.GetCart(), nil)

		got, err := test.unit.AddToCart(test.ctx, &prototransaction.AddToCartReq{
			ProductId:  req.GetProduct().GetId(),
			UserId:     req.GetUserId(),
			QuantityKg: req.GetProduct().GetAmountKg(),
		})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func TestCartService_UpdateCart(t *testing.T) {
	test := newCartTest()
	t.Run("given nil param, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.UpdateCart(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Update(entity.Cart{
			ID:       req.GetId(),
			AmountKG: req.GetProduct().GetAmountKg(),
			UserID:   req.GetUserId(),
		}).Return(entity.Cart{}, errors.New("error"))

		got, err := test.unit.UpdateCart(test.ctx, &prototransaction.UpdateCartReq{
			Id:         req.GetId(),
			QuantityKg: req.GetProduct().GetAmountKg(),
			UserId:     req.GetUserId(),
		})
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Update(entity.Cart{
			ID:       req.GetId(),
			AmountKG: req.GetProduct().GetAmountKg(),
			UserID:   req.GetUserId(),
		}).Return(test.GetCart(), nil)

		got, err := test.unit.UpdateCart(test.ctx, &prototransaction.UpdateCartReq{
			Id:         req.GetId(),
			QuantityKg: req.GetProduct().GetAmountKg(),
			UserId:     req.GetUserId(),
		})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func TestCartService_RemoveCart(t *testing.T) {
	test := newCartTest()
	t.Run("given nil param, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.RemoveCart(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Remove(req.GetId(), req.GetUserId()).Return(errors.New("error"))

		got, err := test.unit.RemoveCart(test.ctx, &prototransaction.RemoveCartReq{
			Id:     req.GetId(),
			UserId: req.GetUserId(),
		})
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetCartProto()

		test.uc.EXPECT().Remove(req.GetId(), req.GetUserId()).Return(nil)

		got, err := test.unit.RemoveCart(test.ctx, &prototransaction.RemoveCartReq{
			Id:     req.GetId(),
			UserId: req.GetUserId(),
		})
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
func (ct *cartTest) Begin(t *testing.T) {
	ct.ctrl = gomock.NewController(t)
	ct.uc = ucmock.NewMockCartUC(ct.ctrl)
	ct.unit = NewCartService(ct.uc)
	ct.ctx = context.Background()
}

func (ct *cartTest) Finish() {
	ct.ctrl.Finish()
}

func (ct *cartTest) GetListReq() *prototransaction.ListCartItemsReq {
	return &prototransaction.ListCartItemsReq{
		UserId: 1,
	}
}
func (ct *cartTest) GetCartProto() *prototransaction.Cart {
	return &prototransaction.Cart{
		Id:     123,
		UserId: 34,
		Product: &prototransaction.Product{
			Id:         1,
			ShopId:     2,
			Name:       "asdf",
			AmountKg:   123,
			PricePerKg: 2,
			TotalPrice: 246,
		},
	}
}
func (ct *cartTest) GetCart() entity.Cart {
	return entity.Cart{
		ID:     12,
		UserID: 34,
		Product: entity.Product{
			ID:         1,
			ShopID:     2,
			Name:       "asdf",
			AmountKG:   123,
			PricePerKG: 2,
			TotalPrice: 246,
		},
	}
}
func (ct *cartTest) GetCarts() []entity.Cart {
	return []entity.Cart{
		{
			ID:     15,
			UserID: 36,
			Product: entity.Product{
				ID:         2,
				ShopID:     1,
				Name:       "asdf",
				AmountKG:   111,
				PricePerKG: 3,
				TotalPrice: 333,
			},
		},
		{
			ID:     12,
			UserID: 34,
			Product: entity.Product{
				ID:         1,
				ShopID:     2,
				Name:       "asdf",
				AmountKG:   123,
				PricePerKG: 2,
				TotalPrice: 246,
			},
		},
	}
}
