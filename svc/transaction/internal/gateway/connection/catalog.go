package connection

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"google.golang.org/grpc"
)

type CatalogClient struct {
	conn *grpc.ClientConn
	catalogproto.MarketplaceCatalogClient
}

func NewCatalogClient(conn *grpc.ClientConn) *CatalogClient {
	client := catalogproto.NewMarketplaceCatalogClient(conn)
	return &CatalogClient{
		conn:                     conn,
		MarketplaceCatalogClient: client,
	}
}

func (cc *CatalogClient) Close() error {
	return cc.conn.Close()
}
