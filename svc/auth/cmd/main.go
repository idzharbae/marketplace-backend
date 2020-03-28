package main

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/delivery/grpc"
	"log"
	"path/filepath"
	"runtime"

	"github.com/idzharbae/marketplace-backend/svc/auth/internal/app"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	// init app
	marketplaceApp, err := app.NewAuth(basepath + "/../config/auth.json")
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
