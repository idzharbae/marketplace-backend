package app

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase"
)

type UseCases struct {
	UserUC    internal.UserUC
	TokenUC   internal.TokenUC
	PaymentUC internal.PaymentUC
}

func NewUsecases(cfg config.Config, repos *Repos) *UseCases {
	return &UseCases{
		UserUC:    usecase.NewUser(repos.UserReader, repos.UserWriter),
		TokenUC:   usecase.NewToken(cfg),
		PaymentUC: usecase.NewPaymentUC(repos.UserWriter),
	}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}
