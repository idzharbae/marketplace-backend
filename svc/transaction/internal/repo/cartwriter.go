package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
)

type CartWriter struct {
	db connection.Gormw
}

func NewCartWriter(db connection.Gormw) *CartWriter {
	return &CartWriter{db: db}
}

func (cw *CartWriter) Create(cart entity.Cart) (entity.Cart, error) {
	return entity.Cart{}, nil
}
func (cw *CartWriter) Update(cart entity.Cart) (entity.Cart, error) {
	return entity.Cart{}, nil
}
func (cw *CartWriter) DeleteByID(cartID int64) error {
	return nil
}
