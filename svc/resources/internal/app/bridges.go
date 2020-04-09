package app

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/bridge/fileio"
)

type Bridges struct {
	FileIO internal.FileIO
}

func NewBridges() *Bridges {
	return &Bridges{
		FileIO: fileio.NewIO(),
	}
}
