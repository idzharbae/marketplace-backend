package app

import (
	"github.com/idzharbae/marketplace-backend/internal"
	"github.com/idzharbae/marketplace-backend/internal/usecase"
)

type UseCases struct {
	ProductUC internal.ProductUC
}

func NewUsecase(repos *Repos) *UseCases {
	return &UseCases{
		usecase.NewProductUC(repos.ProductReader),
	}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}
