package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
)

type paymentTest struct {
	ctrl *gomock.Controller
	uc   *ucmock.MockPaymentUC
	unit *PaymentService
}

func newPaymentTest() *paymentTest {
	return &paymentTest{}
}

func (pt *paymentTest) Begin(t *testing.T) {
	pt.ctrl = gomock.NewController(t)
	pt.uc = ucmock.NewMockPaymentUC(pt.ctrl)
	pt.unit = NewPaymentService(pt.uc)
}

func (pt *paymentTest) Finish() {
	pt.ctrl.Finish()
}

func TestPaymentService_TopUp(t *testing.T) {
	test := newPaymentTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.TopUp(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.TopUpReq{
			UserId: 12,
			Amount: 1337,
		}
		test.uc.EXPECT().TopUp(request.TopUp{
			UserID: req.GetUserId(),
			Amount: req.GetAmount(),
		}).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.TopUp(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.TopUpReq{
			UserId: 12,
			Amount: 1337,
		}
		resp := entity.User{ID: req.GetUserId(), Saldo: 1000 + req.GetAmount()}
		test.uc.EXPECT().TopUp(request.TopUp{
			UserID: req.GetUserId(),
			Amount: req.GetAmount(),
		}).Return(resp, nil)

		got, err := test.unit.TopUp(context.Background(), req)
		assert.NotNil(t, got)
		assert.Equal(t, got.GetSaldo(), resp.Saldo)
		assert.Nil(t, err)
	})
}

func TestPaymentService_TransferSaldo(t *testing.T) {
	test := newPaymentTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.TransferSaldo(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.TransferSaldoReq{
			SenderId:       23,
			ReceiverId:     24,
			TransferAmount: 1337,
		}
		test.uc.EXPECT().Transfer(request.Transfer{
			SenderID:       req.GetSenderId(),
			ReceiverID:     req.GetReceiverId(),
			TransferAmount: req.GetTransferAmount(),
		}).Return(authproto.TransferSaldoResp{}, errors.New("error"))

		got, err := test.unit.TransferSaldo(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.TransferSaldoReq{
			SenderId:       23,
			ReceiverId:     24,
			TransferAmount: 1337,
		}
		resp := authproto.TransferSaldoResp{
			SenderId:      23,
			ReceiverId:    24,
			SenderSaldo:   1231,
			ReceiverSaldo: 43434,
		}
		test.uc.EXPECT().Transfer(request.Transfer{
			SenderID:       req.GetSenderId(),
			ReceiverID:     req.GetReceiverId(),
			TransferAmount: req.GetTransferAmount(),
		}).Return(resp, nil)

		got, err := test.unit.TransferSaldo(context.Background(), req)
		assert.NotNil(t, got)
		assert.Equal(t, got.GetReceiverId(), resp.ReceiverId)
		assert.Equal(t, got.GetSenderId(), resp.SenderId)
		assert.Equal(t, got.GetReceiverSaldo(), resp.ReceiverSaldo)
		assert.Equal(t, got.GetSenderSaldo(), resp.SenderSaldo)
		assert.Nil(t, err)
	})
}
