package internal

import (
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

//go:generate mockgen -destination=repo/repomock/productreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/internal ProductReader
type ProductReader interface {
	List(req requests.ListProduct) ([]entity.Product, error)
	GetByID(productID int32) (entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)
}

//go:generate mockgen -destination=repo/repomock/productwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/internal ProductWriter
type ProductWriter interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(productID int32) error
}
