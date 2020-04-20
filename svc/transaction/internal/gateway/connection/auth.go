package connection

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"google.golang.org/grpc"
)

type AuthClient struct {
	conn *grpc.ClientConn
	authproto.MarketplaceAuthClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	client := authproto.NewMarketplaceAuthClient(conn)
	return &AuthClient{
		conn:                  conn,
		MarketplaceAuthClient: client,
	}
}

func (cc *AuthClient) Close() error {
	return cc.conn.Close()
}
