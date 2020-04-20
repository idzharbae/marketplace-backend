package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

type OrderWriter struct {
	db connection.Gormw
}

func NewOrderWriter(db connection.Gormw) *OrderWriter {
	return &OrderWriter{db: db}
}

func (ow *OrderWriter) CreateFromCartsAndSubstractCustomerSaldo(req request.CheckoutReq) (entity.Order, error) {
	return entity.Order{}, nil
}
func (ow *OrderWriter) UpdateOrderStatusAndAddShopSaldo(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}
