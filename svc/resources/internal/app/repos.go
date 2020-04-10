package app

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/connection"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/storage"
)

type Repos struct {
	FileWriter      internal.FileWriter
	OwnershipWriter internal.OwnershipWriter
	OwnershipReader internal.OwnershipReader
}

func NewRepos(bridges *Bridges, cfg config.Config) (*Repos, error) {
	masterConn, err := connection.NewConnection(cfg.Db.DBEngine, connection.GetConnectionParams(cfg.Db.Master))
	slaveConn, err := connection.NewConnection(cfg.Db.DBEngine, connection.GetConnectionParams(cfg.Db.Slave))
	if err != nil {
		return nil, err
	}
	return &Repos{
		FileWriter:      storage.NewFileWriter(bridges.FileIO, cfg),
		OwnershipWriter: db.NewOwnershipWriter(masterConn),
		OwnershipReader: db.NewOwnershipReader(slaveConn),
	}, nil
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
