package service

import "github.com/idzharbae/marketplace-backend/internal/app"

type ProductService struct {

}

func NewProductService(a *app.Marketplace) *ProductService {
	return &ProductService{}
}
