package service

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/app"
)

type Services struct {
	*TransactionService
}

func GetServices(a *app.Transaction) *Services {
	return &Services{
		TransactionService: NewTransaction(),
	}
}
