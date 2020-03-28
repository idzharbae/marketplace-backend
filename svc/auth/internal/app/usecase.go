package app

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase"
)

type UseCases struct {
	UserUC  internal.UserUC
	TokenUC internal.TokenUC
}

func NewUsecases(repos *Repos) *UseCases {
	return &UseCases{
		UserUC:  usecase.NewUser(repos.UserReader, repos.UserWriter),
		TokenUC: usecase.NewToken(repos.UserReader),
	}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}