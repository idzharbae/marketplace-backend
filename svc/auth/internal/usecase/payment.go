package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/constant"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type Payment struct {
	userWriter         internal.UserWriter
	userReader         internal.UserReader
	saldoHistoryWriter internal.SaldoHistoryWriter
}

func NewPaymentUC(reader internal.UserReader, writer internal.UserWriter, shw internal.SaldoHistoryWriter) *Payment {
	return &Payment{userWriter: writer, saldoHistoryWriter: shw, userReader: reader}
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
	user, err := p.userReader.GetByID(req.UserID)
	if err != nil {
		return entity.User{}, err
	}
	req.UserType = user.Type
	desc := p.getPaymentDescription(req)

	p.saldoHistoryWriter.Create(entity.SaldoHistory{
		UserID:       req.UserID,
		Description:  desc,
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

func (p *Payment) getPaymentDescription(req request.TopUp) string {
	desc := ""
	if req.UserType == constant.UserBuyerType {
		if req.Amount > 0 {
			desc = "payment_refund"
		} else {
			desc = "payment"
		}
	} else {
		if req.Amount > 0 {
			desc = "received_payment"
		} else {
			desc = "payment_refund"
		}
	}
	return desc
}
