package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
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
	ot.unit = NewOrder(ot.reader, ot.writer, ot.cartReader)
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
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return(nil, errors.New("error"))

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("payment not equal price, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 2, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 3, Product: entity.Product{PricePerKG: 1000, AmountKG: 2}},
		}, nil)

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("error when creating order, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}

		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return([]entity.Cart{
			{ID: 1, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 2, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 3, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
		}, nil)
		test.writer.EXPECT().CreateFromCartsAndSubstractCustomerSaldo(req).Return(entity.Order{}, errors.New("error"))

		got, err := test.unit.CreateFromCarts(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Order{}, got)
	})
	t.Run("create order success, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.CheckoutReq{
			CartIDs:       []int64{1, 23, 4},
			PaymentAmount: 3000,
		}
		respCart := []entity.Cart{
			{ID: 1, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 2, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
			{ID: 3, Product: entity.Product{PricePerKG: 1000, AmountKG: 1}},
		}
		respOrder := entity.Order{
			ID:         1,
			UserID:     2,
			Products:   []entity.Product{respCart[0].Product, respCart[1].Product, respCart[2].Product},
			TotalPrice: req.PaymentAmount,
			Status:     1,
		}
		test.cartReader.EXPECT().GetByIDs(req.CartIDs).Return(respCart, nil)
		test.writer.EXPECT().CreateFromCartsAndSubstractCustomerSaldo(req).Return(respOrder, nil)

		got, err := test.unit.CreateFromCarts(req)
		assert.Nil(t, err)
		assert.Equal(t, respOrder, got)
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
