package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
)

type CartReader struct {
	db connection.Gormw
}

func NewCartReader(db connection.Gormw) *CartReader {
	return &CartReader{db: db}
}

func (cr *CartReader) ListByUserID(userID int64) ([]entity.Cart, error) {
	var carts []model.Cart
	err := cr.db.Where("user_id=?", userID).Find(&carts).Error()
	if err != nil {
		return nil, err
	}
	return converter.CartModelsToEntities(carts), nil
}
