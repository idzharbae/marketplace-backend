package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type ReviewReader struct {
	db connection.Gormw
}

func NewReviewReader(db connection.Gormw) *ReviewReader {
	return &ReviewReader{db: db}
}

func (r *ReviewReader) ListByProductID(productID int64, pagination requests.Pagination) ([]entity.Review, error) {
	var reviews []model.Review
	db := r.db.Where("product_id=?", productID).Order("created_at DESC")
	if pagination.Limit > 0 {
		db = db.Limit(pagination.Limit)
	}
	if pagination.Page > 0 {
		db = db.Offset(pagination.OffsetFromPagination())
	}
	err := db.Find(&reviews).Error()
	if err != nil {
		return nil, err
	}
	return converter.ReviewModelsToEntities(reviews), nil
}
func (r *ReviewReader) ListByShopID(shopID int64, pagination requests.Pagination) ([]entity.Review, error) {
	var reviews []model.Review
	db := r.db.Where("shop_id=?", shopID).Order("created_at DESC")
	if pagination.Limit > 0 {
		db = db.Limit(pagination.Limit)
	}
	if pagination.Page > 0 {
		db = db.Offset(pagination.OffsetFromPagination())
	}
	err := db.Find(&reviews).Error()
	if err != nil {
		return nil, err
	}
	return converter.ReviewModelsToEntities(reviews), nil
}

func (r *ReviewReader) GetByID(reviewID int64) (entity.Review, error) {
	var review model.Review
	err := r.db.Where("id=?", reviewID).First(&review).Error()
	if err != nil {
		return entity.Review{}, err
	}
	return converter.ReviewModelToEntity(review), nil
}
