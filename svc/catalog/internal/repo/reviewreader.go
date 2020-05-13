package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
)

type ReviewReader struct {
	db connection.Gormw
}

func NewReviewReader(db connection.Gormw) *ReviewReader {
	return &ReviewReader{db: db}
}

func (r *ReviewReader) ListByProductID(productID int64) ([]entity.Review, error) {
	return nil, nil
}
func (r *ReviewReader) ListByShopID(shopID int64) ([]entity.Review, error) {
	return nil, nil
}

func (r *ReviewReader) GetByID(reviewID int64) (entity.Review, error) {
	return entity.Review{}, nil
}
