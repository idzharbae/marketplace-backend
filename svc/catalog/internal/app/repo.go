package app

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
)

type Repos struct {
	ProductReader internal.ProductReader
	ProductWriter internal.ProductWriter
}

func NewRepos(cfg config.Config) *Repos {
	connMaster, err := getMasterConn(cfg)
	if err != nil {
		panic(err)
	}
	connSlave, err := getSlaveConn(cfg)
	if err != nil {
		panic(err)
	}
	productReader := repo.NewProductReader(connSlave)
	productWriter := repo.NewProductWriter(connMaster)
	return &Repos{
		ProductReader: productReader,
		ProductWriter: productWriter,
	}
}

func getSlaveConn(cfg config.Config) (connection.Gormw, error) {
	connSlave, err := connection.NewConnection(cfg.Db.DBEngine, connection.GetConnectionParams(cfg.Db.Master))
	if err != nil {
		return nil, err
	}
	if cfg.Db.Debug {
		return connSlave.Debug(), nil
	}
	return connSlave, nil
}

func getMasterConn(cfg config.Config) (connection.Gormw, error) {
	connMaster, err := connection.NewConnection(cfg.Db.DBEngine, connection.GetConnectionParams(cfg.Db.Master))
	if err != nil {
		return nil, err
	}
	if cfg.Db.Debug {
		return connMaster.Debug(), nil
	}
	return connMaster, nil
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
