package app

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/usecase"
)

type UseCases struct {
	CartUC internal.CartUC
}

func NewUseCases(repos *Repos) *UseCases {
	return &UseCases{
		CartUC: usecase.NewCart(repos.CartReader, repos.CartWriter),
	}
}
