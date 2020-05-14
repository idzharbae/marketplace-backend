package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
)

type ReviewWriter struct {
	db connection.Gormw
}

func NewReviewWriter(db connection.Gormw) *ReviewWriter {
	return &ReviewWriter{db: db}
}

func (w *ReviewWriter) Create(review entity.Review) (entity.Review, error) {
	reviewModel := converter.ReviewEntityToModel(review)
	err := w.db.Save(&reviewModel).Error()
	if err != nil {
		return entity.Review{}, err
	}
	return converter.ReviewModelToEntity(reviewModel), nil
}
func (w *ReviewWriter) Update(review entity.Review) (entity.Review, error) {
	var found model.Review
	err := w.db.Where("id=?", review.ID).First(&found).Error()
	if err != nil {
		return entity.Review{}, err
	}
	savedModel := converter.ReviewEntityToModel(review)
	err = w.db.Save(&savedModel).Error()
	if err != nil {
		return entity.Review{}, err
	}
	return converter.ReviewModelToEntity(savedModel), nil
}
func (w *ReviewWriter) Delete(review entity.Review) error {
	return w.db.Delete(&model.Review{ID: review.ID}).Error()
}
