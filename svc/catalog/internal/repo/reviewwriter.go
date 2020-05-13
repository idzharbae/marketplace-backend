package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
)

type ReviewWriter struct {
	db connection.Gormw
}

func NewReviewWriter(db connection.Gormw) *ReviewWriter {
	return &ReviewWriter{db: db}
}

func (w *ReviewWriter) Create(review entity.Review) (entity.Review, error) {
	return entity.Review{}, nil
}
func (w *ReviewWriter) Update(review entity.Review) (entity.Review, error) {
	return entity.Review{}, nil
}
func (w *ReviewWriter) Delete(review entity.Review) error {
	return nil
}
