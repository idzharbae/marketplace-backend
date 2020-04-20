package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type PaymentService struct {
	PaymentUC internal.PaymentUC
}

func NewPaymentService(paymentUC internal.PaymentUC) *PaymentService {
	return &PaymentService{PaymentUC: paymentUC}
}

func (p *PaymentService) TopUp(ctx context.Context, in *authproto.TopUpReq) (*authproto.TopUpResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := p.PaymentUC.TopUp(request.TopUp{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &authproto.TopUpResp{
		UserId: res.ID,
		Saldo:  res.Saldo,
	}, nil
}

func (p *PaymentService) UpdateSaldo(ctx context.Context, in *authproto.TopUpReq) (*authproto.TopUpResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := p.PaymentUC.UpdateSaldo(request.TopUp{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &authproto.TopUpResp{
		UserId: res.ID,
		Saldo:  res.Saldo,
	}, nil
}

func (p *PaymentService) TransferSaldo(ctx context.Context, in *authproto.TransferSaldoReq) (*authproto.TransferSaldoResp, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := p.PaymentUC.Transfer(request.Transfer{
		SenderID:       in.GetSenderId(),
		ReceiverID:     in.GetReceiverId(),
		TransferAmount: in.GetTransferAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}
