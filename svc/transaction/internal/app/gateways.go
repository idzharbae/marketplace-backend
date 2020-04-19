package app

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection"
	"google.golang.org/grpc"
)

type Gateways struct {
	CatalogGateway internal.CatalogGateway
}

func NewGateways(cfg config.Config) (*Gateways, error) {
	catalogConn, err := grpc.Dial(cfg.Gateways.Catalog.Grpc.Port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	catalogClient := connection.NewCatalogClient(catalogConn)

	return &Gateways{CatalogGateway: gateway.NewCatalog(catalogClient)}, nil
}
