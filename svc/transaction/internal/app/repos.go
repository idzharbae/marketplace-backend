package app

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/repo/connection"
)

type Repos struct {
	CartReader  internal.CartReader
	CartWriter  internal.CartWriter
	OrderReader internal.OrderReader
	OrderWriter internal.OrderWriter

	connMaster connection.Gormw
	connSlave  connection.Gormw
}

func NewRepos(cfg config.Config, gateways *Gateways) *Repos {
	connMaster, err := getMasterConn(cfg)
	if err != nil {
		panic(err)
	}
	connSlave, err := getSlaveConn(cfg)
	if err != nil {
		panic(err)
	}

	cartReader := repo.NewCartReader(connSlave)
	cartWriter := repo.NewCartWriter(connMaster)
	orderReader := repo.NewOrderReader(connSlave)
	orderWriter := repo.NewOrderWriter(connMaster, gateways.AuthGateway, gateways.CatalogGateway)

	return &Repos{
		CartWriter:  cartWriter,
		CartReader:  cartReader,
		OrderReader: orderReader,
		OrderWriter: orderWriter,
		connSlave:   connSlave,
		connMaster:  connMaster,
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
	err := r.connMaster.Close()
	if err != nil {
		errs = append(errs, err)
	}
	err = r.connSlave.Close()
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}
