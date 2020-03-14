package app

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/usecase"
)

type UseCases struct {
	ProductUC internal.ProductUC
	ShopUC    internal.ShopUC
}

func NewUsecase(repos *Repos) *UseCases {
	return &UseCases{
		ProductUC: usecase.NewProduct(repos.ProductReader, repos.ProductWriter),
		ShopUC:    usecase.NewShop(repos.ShopReader, repos.ShopWriter),
	}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}
