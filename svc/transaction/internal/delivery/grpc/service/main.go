package service

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/app"
)

type Services struct {
	*CartService
	*OrderService
}

func GetServices(a *app.Transaction) *Services {
	return &Services{
		CartService:  NewCartService(a.UseCases.CartUC),
		OrderService: NewOrderService(a.UseCases.OrderUC),
	}
}
