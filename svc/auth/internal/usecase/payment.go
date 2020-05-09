package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type Payment struct {
	userWriter         internal.UserWriter
	saldoHistoryWriter internal.SaldoHistoryWriter
}

func NewPaymentUC(writer internal.UserWriter, shw internal.SaldoHistoryWriter) *Payment {
	return &Payment{userWriter: writer, saldoHistoryWriter: shw}
}

// exposed to graphql so need validation
func (p *Payment) TopUp(req request.TopUp) (entity.User, error) {
	if req.Amount < 0 {
		return entity.User{}, errors.New("topup amount cant be negative")
	}
	res, err := p.userWriter.UpdateSaldo(req)
	if err != nil {
		return entity.User{}, err
	}
	p.saldoHistoryWriter.Create(entity.SaldoHistory{
		UserID:       req.UserID,
		Description:  "topup",
		ChangeAmount: req.Amount,
	})
	return res, nil
}

func (p *Payment) UpdateSaldo(req request.TopUp) (entity.User, error) {
	p.saldoHistoryWriter.Create(entity.SaldoHistory{
		UserID:       req.UserID,
		Description:  "payment",
		ChangeAmount: req.Amount,
	})
	return p.userWriter.UpdateSaldo(req)
}

func (p *Payment) Transfer(req request.Transfer) (authproto.TransferSaldoResp, error) {
	if req.TransferAmount < 0 {
		return authproto.TransferSaldoResp{}, errors.New("transfer amount cant be negative")
	}
	return p.userWriter.TransferSaldo(req)
}

func (p *Payment) saveSaldoHistory(userID, amount int64) error {
	_, err := p.saldoHistoryWriter.Create(entity.SaldoHistory{
		UserID:       userID,
		Description:  "topup",
		ChangeAmount: amount,
	})
	return err
}
