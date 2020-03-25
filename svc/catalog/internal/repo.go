package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

//go:generate mockgen -destination=repo/repomock/productreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductReader
type ProductReader interface {
	ListAll(pagination requests.Pagination) ([]entity.Product, error)
	ListByShopID(shopID int64, pagination requests.Pagination) ([]entity.Product, error)

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

//go:generate mockgen -destination=repo/repomock/shopreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopReader
type ShopReader interface {
	ListAll(pagination requests.Pagination) ([]entity.Shop, error)

	GetByID(shopID int32) (entity.Shop, error)
	GetBySlug(slug string) (entity.Shop, error)
}

//go:generate mockgen -destination=repo/repomock/shopwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopWriter
type ShopWriter interface {
	Create(shop entity.Shop) (entity.Shop, error)
	Update(shop entity.Shop) (entity.Shop, error)
	DeleteByID(shopID int32) error
	DeleteBySlug(shopSlug string) error
}
