package service

import "github.com/idzharbae/marketplace-backend/internal/app"

// Services list all service struct
type Services struct {
	*ProductService
}

func GetServices(a *app.Marketplace) *Services {
	return &Services{
		ProductService: NewProductService(a.UseCases.ProductUC),
	}
}
