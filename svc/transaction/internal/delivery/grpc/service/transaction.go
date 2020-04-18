package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type TransactionService struct {
}

func NewTransaction() *TransactionService {
	return &TransactionService{}
}

func (ts *TransactionService) AddToCart(context.Context, *prototransaction.AddToCartReq) (*prototransaction.AddToCartResp, error) {
	panic("implement me")
}

func (ts *TransactionService) Checkout(context.Context, *prototransaction.CheckoutReq) (*prototransaction.CheckoutResp, error) {
	panic("implement me")
}

func (ts *TransactionService) Fulfill(context.Context, *prototransaction.FulfillReq) (*prototransaction.FulfillResp, error) {
	panic("implement me")
}
