package service

import "github.com/idzharbae/marketplace-backend/svc/catalog/internal/app"

// Services list all service struct
type Services struct {
	*ProductService
	*ReviewService
}

func GetServices(a *app.Marketplace) *Services {
	return &Services{
		ProductService: NewProductService(a.UseCases.ProductUC),
		ReviewService:  NewReviewService(a.UseCases.ReviewUC),
	}
}
