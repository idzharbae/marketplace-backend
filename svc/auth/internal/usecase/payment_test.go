package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/repomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

type paymentTest struct {
	ctrl *gomock.Controller
	repo *repomock.MockUserWriter
	shw  *repomock.MockSaldoHistoryWriter
	unit internal.PaymentUC
}

func newPaymentTest() *paymentTest {
	return &paymentTest{}
}
func (pt *paymentTest) Begin(t *testing.T) {
	pt.ctrl = gomock.NewController(t)
	pt.repo = repomock.NewMockUserWriter(pt.ctrl)
	pt.shw = repomock.NewMockSaldoHistoryWriter(pt.ctrl)
	pt.unit = NewPaymentUC(pt.repo, pt.shw)
}

func (pt *paymentTest) Finish() {
	pt.ctrl.Finish()
}

func TestPayment_TopUp(t *testing.T) {
	test := newPaymentTest()
	t.Run("given negative amount, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.TopUp{
			UserID: 123,
			Amount: -12331,
		}
		got, err := test.unit.TopUp(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("error when updating saldo, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := request.TopUp{
			UserID: 123,
			Amount: 1337,
		}
		test.repo.EXPECT().UpdateSaldo(req).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.TopUp(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("no error, save history saldo and return user entity", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		req := request.TopUp{
			UserID: 123,
			Amount: 1337,
		}
		test.repo.EXPECT().UpdateSaldo(req).Return(entity.User{ID: req.UserID, Saldo: req.Amount}, nil)
		test.shw.EXPECT().Create(entity.SaldoHistory{
			UserID:       req.UserID,
			ChangeAmount: req.Amount,
			Description:  "topup",
		}).Return(entity.SaldoHistory{}, nil)

		got, err := test.unit.TopUp(req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.User{}, got)
	})
}

func TestPayment_Transfer(t *testing.T) {
	test := newPaymentTest()
	t.Run("given negative amount, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       123,
			ReceiverID:     1233,
			TransferAmount: -1231234,
		}
		got, err := test.unit.Transfer(req)
		assert.NotNil(t, err)
		assert.Equal(t, authproto.TransferSaldoResp{}, got)
	})
}
