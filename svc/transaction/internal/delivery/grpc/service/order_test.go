package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/usecase/ucmock"
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

func (ot *orderTest) Finsih() {
	ot.ctrl.Finish()
}

func TestOrderService_Checkout(t *testing.T) {
	test := newOrderTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finsih()

		got, err := test.unit.Checkout(test.ctx, nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
}
