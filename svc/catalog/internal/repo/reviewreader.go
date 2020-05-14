package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
)

type ReviewReader struct {
	db connection.Gormw
}

func NewReviewReader(db connection.Gormw) *ReviewReader {
	return &ReviewReader{db: db}
}

func (r *ReviewReader) ListByProductID(productID int64) ([]entity.Review, error) {
	var reviews []model.Review
	err := r.db.Where("product_id=?", productID).Find(&reviews).Error()
	if err != nil {
		return nil, err
	}
	return converter.ReviewModelsToEntities(reviews), nil
}
func (r *ReviewReader) ListByShopID(shopID int64) ([]entity.Review, error) {
	var reviews []model.Review
	err := r.db.Where("shop_id=?", shopID).Find(&reviews).Error()
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
