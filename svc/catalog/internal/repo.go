package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/lib/pq"
)

//go:generate mockgen -destination=repo/repomock/productreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductReader
type ProductReader interface {
	ListAll(req requests.ListProduct) ([]entity.Product, error)
	ListByShopID(shopID int64, pagination requests.Pagination) ([]entity.Product, error)
	ListByIDs(productIDs pq.Int64Array) ([]entity.Product, error)

	GetByID(productID int32) (entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)
}

//go:generate mockgen -destination=repo/repomock/productwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductWriter
type ProductWriter interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	DeleteByID(productID int32) error
	DeleteBySlug(productSlug string) error
}
