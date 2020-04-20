package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/gatewaymock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

type orderTest struct {
	ctrl       *gomock.Controller
	reader     *repomock.MockOrderReader
	writer     *repomock.MockOrderWriter
	cartReader *repomock.MockCartReader
	catalog    *gatewaymock.MockCatalogGateway
	unit       internal.OrderUC
}

func newOrderTest() *orderTest {
	return &orderTest{}
}

func (ot *orderTest) Begin(t *testing.T) {
	ot.ctrl = gomock.NewController(t)
	ot.reader = repomock.NewMockOrderReader(ot.ctrl)
	ot.writer = repomock.NewMockOrderWriter(ot.ctrl)
	ot.cartReader = repomock.NewMockCartReader(ot.ctrl)
	ot.catalog = gatewaymock.NewMockCatalogGateway(ot.ctrl)
	ot.unit = NewOrder(ot.reader, ot.writer, ot.cartReader, ot.catalog)
}
func (ot *orderTest) Finish() {
	ot.ctrl.Finish()
}

func TestOrder_CreateFromCarts(t *testing.T) {
	test := newOrderTest()
	t.Run("error when finding carts, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return(nil, errors.New("error"))

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("cart not owned by user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{ID: 2}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{ID: 3}, AmountKG: 1, UserID: 1},
			{ID: 4, Product: entity.Product{ID: 4}, AmountKG: 1, UserID: 69},
		}, nil)

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when finding cart product, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{ID: 2}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{ID: 3}, AmountKG: 1, UserID: 69},
			{ID: 4, Product: entity.Product{ID: 4}, AmountKG: 1, UserID: 69},
		}, nil)
		test.catalog.EXPECT().GetProductsByID([]int64{2, 3, 4}).Return(nil, errors.New("error"))

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("payment not equal price, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{ID: 2}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{ID: 3}, AmountKG: 1, UserID: 69},
			{ID: 4, Product: entity.Product{ID: 4}, AmountKG: 1, UserID: 69},
		}, nil)
		test.catalog.EXPECT().GetProductsByID([]int64{2, 3, 4}).Return(test.GetProducts(), nil)

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error when creating order, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 4000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{ID: 2}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{ID: 3}, AmountKG: 1, UserID: 69},
			{ID: 4, Product: entity.Product{ID: 4}, AmountKG: 1, UserID: 69},
		}, nil)
		test.catalog.EXPECT().GetProductsByID([]int64{2, 3, 4}).Return(test.GetProducts(), nil)
		test.writer.EXPECT().CreateFromCartsAndSubstractCustomerSaldo(test.GetOrderReq()).Return(nil, errors.New("error"))

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("create order success, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			UserID:        69,
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 4000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{ID: 2}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{ID: 3}, AmountKG: 1, UserID: 69},
			{ID: 4, Product: entity.Product{ID: 4}, AmountKG: 1, UserID: 69},
		}, nil)
		test.catalog.EXPECT().GetProductsByID([]int64{2, 3, 4}).Return(test.GetProducts(), nil)
		test.writer.EXPECT().CreateFromCartsAndSubstractCustomerSaldo(test.GetOrderReq()).Return([]entity.Order{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}, nil)

		got, err := test.unit.CreateFromCarts(req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}

func TestOrder_Fulfill(t *testing.T) {
	test := newOrderTest()
	t.Run("given id = 0 should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.Fulfill(entity.Order{ID: 0})
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
}

func (ot *orderTest) GetProducts() []entity.Product {
	return []entity.Product{
		{
			ID:         2,
			ShopID:     2,
			Name:       "asdf",
			PricePerKG: 1000,
			PhotoURL:   "asdf",
		},
		{
			ID:         3,
			ShopID:     2,
			Name:       "asdf",
			PricePerKG: 1000,
			PhotoURL:   "asdf",
		},
		{
			ID:         4,
			ShopID:     2,
			Name:       "asdf",
			PricePerKG: 2000,
			PhotoURL:   "asdf",
		},
	}
}

func (ot *orderTest) GetOrderReq() request.CreateOrderReq {
	return request.CreateOrderReq{
		UserID: 69,
		Carts: []entity.Cart{
			{ID: 1, Product: entity.Product{
				ID:         2,
				ShopID:     2,
				Name:       "asdf",
				PricePerKG: 1000,
				PhotoURL:   "asdf",
			}, AmountKG: 1, UserID: 69},
			{ID: 23, Product: entity.Product{
				ID:         3,
				ShopID:     2,
				Name:       "asdf",
				PricePerKG: 1000,
				PhotoURL:   "asdf",
			}, AmountKG: 1, UserID: 69},
			{ID: 4, Product: entity.Product{
				ID:         4,
				ShopID:     2,
				Name:       "asdf",
				PricePerKG: 2000,
				PhotoURL:   "asdf",
			}, AmountKG: 1, UserID: 69},
		},
		PaymentAmount: 4000,
	}
}
