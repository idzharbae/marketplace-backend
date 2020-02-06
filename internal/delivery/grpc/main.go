package grpc

import (
	"github.com/idzharbae/marketplace-backend/internal/app"
	"github.com/idzharbae/marketplace-backend/internal/delivery/grpc/service"
	pb "github.com/idzharbae/marketplace-backend/marketplaceproto"
	"log"
)

func Start(app *app.Marketplace) {
	// get Service

	// create new service
	var svc pb.MarketplaceServer
	svc = service.GetServices(app)

	// create new grpc server
	grpcServer := NewServer(app.Config.Grpc.Port)

	// register app into delivery (grpc)
	pb.RegisterMarketplaceServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
