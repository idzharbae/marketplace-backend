package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type FulfillmentService struct {
}

func NewFulfillmentService() *FulfillmentService {
	return &FulfillmentService{}
}

func (fs *FulfillmentService) Fulfill(context.Context, *prototransaction.FulfillReq) (*prototransaction.FulfillResp, error) {
	panic("implement me")
}
