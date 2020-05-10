package app

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
)

type Repos struct {
	UserReader         internal.UserReader
	UserWriter         internal.UserWriter
	SaldoHistoryReader internal.SaldoHistoryReader
	SaldoHistoryWriter internal.SaldoHistoryWriter
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

	userReader := repo.NewUserReader(connSlave)
	userWriter := repo.NewUserWriter(connMaster)
	saldoHistoryReader := repo.NewSaldoHistoryReader(connSlave)
	saldoHistoryWriter := repo.NewSaldoHistoryWriter(connMaster)
	return &Repos{
		UserWriter:         userWriter,
		UserReader:         userReader,
		SaldoHistoryWriter: saldoHistoryWriter,
		SaldoHistoryReader: saldoHistoryReader,
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
