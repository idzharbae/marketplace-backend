package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

type Review struct {
	reviewReader internal.ReviewReader
	reviewWriter internal.ReviewWriter
}

func NewReview(reviewReader internal.ReviewReader, reviewWriter internal.ReviewWriter) *Review {
	return &Review{
		reviewReader: reviewReader,
		reviewWriter: reviewWriter,
	}
}

func (r *Review) Get(reveiwID int64) (entity.Review, error) {
	return entity.Review{}, nil
}
func (r *Review) List(req requests.ListReview) ([]entity.Review, error) {
	return nil, nil
}

func (r *Review) Create(review entity.Review) (entity.Review, error) {
	return entity.Review{}, nil
}
func (r *Review) Update(review entity.Review) (entity.Review, error) {
	return entity.Review{}, nil
}
func (r *Review) Delete(review entity.Review) error {
	return nil
}
