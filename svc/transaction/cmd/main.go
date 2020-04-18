package main

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/app"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/delivery/grpc"
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
	transactionApp, err := app.NewTransaction(basepath + "/../config/transaction.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errs := transactionApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	grpc.Start(transactionApp)
}
