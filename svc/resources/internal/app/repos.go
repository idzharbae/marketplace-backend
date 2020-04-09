package app

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/storage"
)

type Repos struct {
	FileWriter internal.FileWriter
}

func NewRepos(bridges *Bridges) *Repos {
	return &Repos{FileWriter: storage.NewFileWriter(bridges.FileIO)}
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
