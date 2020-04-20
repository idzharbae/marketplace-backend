package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
)

type OrderReader struct {
	db connection.Gormw
}

func NewOrderReader(db connection.Gormw) *OrderReader {
	return &OrderReader{db: db}
}

func (or *OrderReader) ListByUserID(userID int64) ([]entity.Order, error) {
	return nil, nil
}
func (or *OrderReader) GetByID(orderID int64) (entity.Order, error) {
	return entity.Order{}, nil
}
