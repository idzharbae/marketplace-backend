package internal

import (
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

type ProductUC interface {
	ListProducts(req requests.ListProduct) ([]entity.Product, error)
}