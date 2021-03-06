package app

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/usecase"
)

type UseCases struct {
	CartUC  internal.CartUC
	OrderUC internal.OrderUC
}

func NewUseCases(repos *Repos, gateways *Gateways) *UseCases {
	return &UseCases{
		CartUC:  usecase.NewCart(repos.CartReader, repos.CartWriter, gateways.CatalogGateway),
		OrderUC: usecase.NewOrder(repos.OrderReader, repos.OrderWriter, repos.CartReader, gateways.CatalogGateway),
	}
}
func (uc *UseCases) Close() []error {
	var errs []error

	return errs
}
