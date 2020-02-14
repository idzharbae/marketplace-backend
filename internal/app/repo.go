package app

import (
	"github.com/idzharbae/marketplace-backend/internal"
	"github.com/idzharbae/marketplace-backend/internal/config"
	"github.com/idzharbae/marketplace-backend/internal/repo"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection"
)

type Repos struct {
	ProductReader internal.ProductReader
}

func NewRepos(cfg config.Config) *Repos {
	connMaster, err := connection.NewConnection(cfg.Db.DBEngine, connection.GetConnectionParams(cfg.Db.Master))
	if err != nil {
		panic(err)
	}
	productReader := repo.NewProductReader(connMaster)
	return &Repos{
		ProductReader: productReader,
	}
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
