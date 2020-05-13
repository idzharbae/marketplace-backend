package app

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/usecase"
)

type UseCases struct {
	ProductUC internal.ProductUC
	ReviewUC  internal.ReviewUC
}

func NewUsecase(repos *Repos) *UseCases {
	return &UseCases{
		ProductUC: usecase.NewProduct(repos.ProductReader, repos.ProductWriter),
		ReviewUC:  usecase.NewReview(repos.ReviewReader, repos.ReviewWriter),
	}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}
