package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type Payment struct {
	userWriter internal.UserWriter
}

func NewPaymentUC(writer internal.UserWriter) *Payment {
	return &Payment{userWriter: writer}
}

func (p *Payment) TopUp(req request.TopUp) (entity.User, error) {
	if req.Amount < 0 {
		return entity.User{}, errors.New("topup amount cant be negative")
	}
	return p.userWriter.UpdateSaldo(req)
}

func (p *Payment) Transfer(req request.Transfer) (authproto.TransferSaldoResp, error) {
	if req.TransferAmount < 0 {
		return authproto.TransferSaldoResp{}, errors.New("topup amount cant be negative")
	}
	return p.userWriter.TransferSaldo(req)
}
