package internal

import (
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

type ProductUC interface {
	List(req requests.ListProduct) ([]entity.Product, error)
	GetByID(productID int32) (entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)

	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(productID int32) error
}
