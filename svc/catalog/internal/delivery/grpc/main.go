package grpc

import (
	pb "github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/delivery/grpc/service"
	"log"
)

func Start(app *app.Marketplace) {
	// get Service

	// create new service
	var svc pb.MarketplaceCatalogServer
	svc = service.GetServices(app)

	// create new grpc server
	grpcServer := NewServer(app.Config.Grpc.Port)

	// register app into delivery (grpc)
	pb.RegisterMarketplaceCatalogServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
