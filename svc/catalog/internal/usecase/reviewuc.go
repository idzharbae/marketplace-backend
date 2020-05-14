package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/constant"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"strconv"
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

func (r *Review) Get(reviewID int64) (entity.Review, error) {
	return r.reviewReader.GetByID(reviewID)
}
func (r *Review) List(req requests.ListReview) ([]entity.Review, error) {
	if req.ProductID != 0 {
		return r.reviewReader.ListByProductID(req.ProductID, req.Pagination)
	}
	return r.reviewReader.ListByShopID(req.ShopID, req.Pagination)
}

func (r *Review) Create(review entity.Review) (entity.Review, error) {
	const op = "ReviewUC::Create()"
	review.ID = 0
	if review.Rating < 0 {
		return entity.Review{}, errors.NewWithPrefix("rating can't be negative", op)
	}
	if review.Rating > constant.MaxRatingValue {
		return entity.Review{}, errors.NewWithPrefix(
			"rating can't be more than "+strconv.FormatInt(int64(constant.MaxRatingValue), 10), op)
	}
	res, err := r.reviewWriter.Create(review)
	if err != nil {
		return entity.Review{}, errors.WithPrefix(err, op)
	}
	return res, nil
}
func (r *Review) Update(review entity.Review) (entity.Review, error) {
	const op = "ReviewUC::Update()"
	if err := review.ValidateRating(); err != nil {
		return entity.Review{}, err
	}
	res, err := r.reviewWriter.Update(review)
	if err != nil {
		return entity.Review{}, errors.WithPrefix(err, op)
	}
	return res, nil
}
func (r *Review) Delete(review entity.Review) error {
	return r.reviewWriter.Delete(review)
}
