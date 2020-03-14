package main

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/delivery/grpc"
	"log"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	// init app
	marketplaceApp, err := app.NewMarketplace(basepath + "/../config/marketplace.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errs := marketplaceApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	grpc.Start(marketplaceApp)
}
