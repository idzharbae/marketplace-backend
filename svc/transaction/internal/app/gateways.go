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
	AuthGateway    internal.AuthGateway

	catalogClient connection.Catalog
	authClient    connection.Auth
}

func NewGateways(cfg config.Config) (*Gateways, error) {
	catalogConn, err := grpc.Dial(cfg.Gateways.Catalog.Grpc.Port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	authConn, err := grpc.Dial(cfg.Gateways.Auth.Grpc.Port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	catalogClient := connection.NewCatalogClient(catalogConn)
	authClient := connection.NewAuthClient(authConn)

	return &Gateways{
		CatalogGateway: gateway.NewCatalog(catalogClient),
		AuthGateway:    gateway.NewAuth(authClient),
		authClient:     authClient,
		catalogClient:  catalogClient,
	}, nil
}

func (g *Gateways) Close() []error {
	var errs []error
	err := g.catalogClient.Close()
	if err != nil {
		errs = append(errs, err)
	}
	err = g.authClient.Close()
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}
