package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type CartReader struct {
	db connection.Gormw
}

func NewCartReader(db connection.Gormw) *CartReader {
	return &CartReader{db: db}
}

func (cr *CartReader) ListByUserID(userID int64, pagination request.Pagination) ([]entity.Cart, error) {
	var carts []model.Cart
	db := cr.db.Where("user_id=?", userID)
	if pagination.Limit > 0 {
		db = db.Limit(pagination.Limit)
	}
	if pagination.Page > 0 {
		db = db.Offset(pagination.OffsetFromPagination())
	}
	err := db.Find(&carts).Error()
	if err != nil {
		return nil, err
	}
	return converter.CartModelsToEntities(carts), nil
}

func (cr *CartReader) GetByIDs(pagination request.Pagination, cartID ...int64) ([]entity.Cart, error) {
	var carts []model.Cart
	db := cr.db.Where("id=ANY(?)", pq.Int64Array(cartID))
	if pagination.Limit > 0 {
		db = db.Limit(pagination.Limit)
	}
	if pagination.Page > 0 {
		db = db.Offset(pagination.OffsetFromPagination())
	}
	err := db.Find(&carts).Error()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return converter.CartModelsToEntities(carts), nil
}
