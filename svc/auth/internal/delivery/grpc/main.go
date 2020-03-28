package grpc

import (
	"log"

	pb "github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/delivery/grpc/service"
)

func Start(app *app.Auth) {
	// get Service

	// create new service
	var svc pb.MarketplaceAuthServer
	svc = service.GetServices(app)

	// create new grpc server
	grpcServer := NewServer(app.Config.Grpc.Port)

	// register app into delivery (grpc)
	pb.RegisterMarketplaceAuthServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
