package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type CheckoutService struct {
}

func NewCheckoutService() *CheckoutService {
	return &CheckoutService{}
}

func (cs *CheckoutService) Checkout(context.Context, *prototransaction.CheckoutReq) (*prototransaction.CheckoutResp, error) {
	panic("implement me")
}
