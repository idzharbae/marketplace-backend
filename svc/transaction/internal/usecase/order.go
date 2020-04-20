package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

type Order struct {
	reader internal.OrderReader
	writer internal.OrderWriter
}

func NewOrder(reader internal.OrderReader, writer internal.OrderWriter) *Order {
	return &Order{writer: writer, reader: reader}
}

func (o *Order) List(req request.ListOrderReq) ([]entity.Order, error) {
	return nil, nil
}

func (o *Order) Get(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o *Order) CreateFromCarts(cartIDs []int64) (entity.Order, error) {
	return entity.Order{}, nil
}
func (o *Order) Fulfill(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}
