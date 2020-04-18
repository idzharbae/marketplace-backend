package usecase

import (
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
	unit internal.PaymentUC
}

func newPaymentTest() *paymentTest {
	return &paymentTest{}
}
func (pt *paymentTest) Begin(t *testing.T) {
	pt.ctrl = gomock.NewController(t)
	pt.repo = repomock.NewMockUserWriter(pt.ctrl)
	pt.unit = NewPaymentUC(pt.repo)
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
