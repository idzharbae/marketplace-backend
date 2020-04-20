package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
)

type OrderWriter struct {
	db connection.Gormw
}

func NewOrderWriter(db connection.Gormw) *OrderWriter {
	return &OrderWriter{db: db}
}

func (ow *OrderWriter) Create(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}
func (ow *OrderWriter) Update(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}
func (ow *OrderWriter) DeleteByID(orderID int64) error {
	return nil
}
