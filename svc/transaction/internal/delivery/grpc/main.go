package grpc

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/delivery/grpc/service"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"

	"log"
)

func Start(app *app.Transaction) {
	// get Service

	// create new service
	var svc prototransaction.MarketplaceTransactionServer
	svc = service.GetServices(app)

	// create new grpc server
	grpcServer := NewServer(app.Config.Grpc.Port)

	// register app into delivery (grpc)
	prototransaction.RegisterMarketplaceTransactionServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
