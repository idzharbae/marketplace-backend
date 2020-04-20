package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/gatewaymock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

type orderWriterTest struct {
	ctrl    *gomock.Controller
	db      *gormmock.MockGormw
	auth    *gatewaymock.MockAuthGateway
	catalog *gatewaymock.MockCatalogGateway
	unit    internal.OrderWriter
}

func newOrderWriterTest() *orderWriterTest {
	return &orderWriterTest{}
}

func TestOrderWriter_CreateFromCartsAndSubstractCustomerSaldo(t *testing.T) {
	test := newOrderWriterTest()
	t.Run("error when fetching user data, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("not enough saldo, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount - 1}, nil)

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when saving order, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount}, nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, int64(1337), (*arg).OrderID)
			(*arg).ID = 123
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		test.db.EXPECT().Rollback()

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when creating payment data, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount}, nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, int64(1337), (*arg).OrderID)
			(*arg).ID = 123
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1338
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		test.db.EXPECT().Rollback()

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when deleting cart data, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount}, nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, (*arg).OrderID, int64(1337))
			(*arg).ID = 123
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1338
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, int64(1338), (*arg).OrderID)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Delete(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		test.db.EXPECT().Rollback()

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when substracting user saldo, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount}, nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, (*arg).OrderID, int64(1337))
			(*arg).ID = 123
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1338
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, int64(1338), (*arg).OrderID)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Delete(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.auth.EXPECT().UpdateUserSaldo(req.UserID, -req.PaymentAmount).Return(entity.User{}, errors.New("error"))
		test.db.EXPECT().Rollback()

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("substract saldo successful, should return order data", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := test.GetReq()
		test.auth.EXPECT().GetUserByID(req.UserID).Return(entity.User{ID: req.UserID, Saldo: req.PaymentAmount}, nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1337
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, (*arg).OrderID, int64(1337))
			(*arg).ID = 123
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Order) *gormmock.MockGormw {
			(*arg).ID = 1338
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.Payment) *gormmock.MockGormw {
			assert.Equal(t, int64(1338), (*arg).OrderID)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Delete(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.auth.EXPECT().UpdateUserSaldo(req.UserID, -req.PaymentAmount).Return(entity.User{}, nil)
		test.db.EXPECT().Commit()

		got, err := test.unit.CreateFromCartsAndSubstractCustomerSaldo(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, 2, len(got))
	})
}

func (ow *orderWriterTest) Begin(t *testing.T) {
	ow.ctrl = gomock.NewController(t)
	ow.db = gormmock.NewMockGormw(ow.ctrl)
	ow.auth = gatewaymock.NewMockAuthGateway(ow.ctrl)
	ow.catalog = gatewaymock.NewMockCatalogGateway(ow.ctrl)
	ow.unit = NewOrderWriter(ow.db, ow.auth, ow.catalog)
}

func (ow *orderWriterTest) Finish() {
	ow.ctrl.Finish()
}

func (ow *orderWriterTest) GetReq() request.CreateOrderReq {
	return request.CreateOrderReq{
		UserID: 1,
		Carts: []entity.Cart{
			{
				ID: 1,
				Product: entity.Product{
					ID:         2,
					ShopID:     1,
					Name:       "asdf",
					AmountKG:   1,
					PricePerKG: 1000,
					TotalPrice: 1000,
					PhotoURL:   "asdf",
				},
				UserID:   1,
				AmountKG: 1,
			},
			{
				ID: 2,
				Product: entity.Product{
					ID:         3,
					ShopID:     1,
					Name:       "asdf",
					AmountKG:   1,
					PricePerKG: 1000,
					TotalPrice: 1000,
					PhotoURL:   "asdf",
				},
				UserID:   1,
				AmountKG: 1,
			},
			{
				ID: 3,
				Product: entity.Product{
					ID:         4,
					ShopID:     2,
					Name:       "asdf",
					AmountKG:   1,
					PricePerKG: 1900,
					TotalPrice: 1900,
					PhotoURL:   "asdf",
				},
				UserID:   1,
				AmountKG: 1,
			},
		},
		PaymentAmount: 3900,
	}
}
