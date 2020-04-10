package app

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/usecase"
)

type Usecases struct {
	File internal.FileUC
}

func NewUsecases(repos *Repos) *Usecases {
	return &Usecases{
		File: usecase.NewFile(repos.FileWriter, repos.OwnershipWriter, repos.OwnershipReader),
	}
}

func (ucs *Usecases) Close() []error {
	var errs []error

	return errs
}
