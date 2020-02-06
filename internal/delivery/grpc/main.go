package grpc

import (
	"github.com/idzharbae/marketplace-backend/internal/app"
	"github.com/idzharbae/marketplace-backend/internal/delivery/grpc/service"
	pb "github.com/idzharbae/marketplace-backend/marketplaceproto"
)

func Start(app *app.Marketplace) {
	var svc pb.MarketplaceServer
	svc = service.GetServices(app)
	grpcServer := NewServer(app.Config.Grpc.Address)
	pb.RegisterMarketplaceServer(grpcServer.Server(), svc)
}