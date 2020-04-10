package grpc

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/delivery/grpc/service"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	"log"
)

func Start(app *app.Resources) {
	// get Service

	// create new service
	var svc protoresources.MarketplaceResourcesServer
	svc = service.GetServices(app)

	// create new grpc server
	grpcServer := NewServer(app.Config.Grpc.Port)

	// register app into delivery (grpc)
	protoresources.RegisterMarketplaceResourcesServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
