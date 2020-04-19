package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
)

type CartReader struct {
	db connection.Gormw
}

func NewCartReader(db connection.Gormw) *CartReader {
	return &CartReader{db: db}
}

func (cr *CartReader) ListByUserID(userID int64) ([]entity.Cart, error) {
	return nil, nil
}
