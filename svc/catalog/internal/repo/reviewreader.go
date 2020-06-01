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

func (r *ReviewReader) GetByCustomerIDAndProductID(customerID, productID int64) (entity.Review, error) {
	var review model.Review
	err := r.db.Where("user_id=?", customerID).Where("product_id=?", productID).First(&review).Error()
	if err != nil {
		return entity.Review{}, err
	}
	return converter.ReviewModelToEntity(review), nil
}

func (r *ReviewReader) GetTotalAndAverageByShopID(shopID int64) (requests.TotalAndAverageReview, error) {
	var res model.Review
	// need better way to store these two values
	err := r.db.Select("COUNT(*) AS shop_id, AVG(rating) AS rating").Where("shop_id=?", shopID).Find(&res).Error()
	if err != nil {
		return requests.TotalAndAverageReview{}, err
	}
	return requests.TotalAndAverageReview{
		Total:   int32(res.ShopID),
		Average: float32(res.Rating),
	}, nil
}
func (r *ReviewReader) GetTotalAndAverageByProductID(productID int64) (requests.TotalAndAverageReview, error) {
	var res model.Review
	err := r.db.Select("COUNT(*) AS shop_id, AVG(rating) AS rating").Where("product_id=?", productID).Find(&res).Error()
	if err != nil {
		return requests.TotalAndAverageReview{}, err
	}
	return requests.TotalAndAverageReview{
		Total:   int32(res.ShopID),
		Average: float32(res.Rating),
	}, nil
}
