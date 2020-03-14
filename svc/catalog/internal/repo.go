package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

//go:generate mockgen -destination=repo/repomock/productreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductReader
type ProductReader interface {
	List(req requests.ListProduct) ([]entity.Product, error)
	GetByID(productID int32) (entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)
	GetByShopID(shopID int32) ([]entity.Product, error)
}

//go:generate mockgen -destination=repo/repomock/productwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductWriter
type ProductWriter interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(productID int32) error
}

//go:generate mockgen -destination=repo/repomock/shopreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopReader
type ShopReader interface {
	List(req requests.ListShop) ([]entity.Shop, error)
	GetByID(shopID int32) (entity.Shop, error)
	GetBySlug(slug string) (entity.Shop, error)
}

//go:generate mockgen -destination=repo/repomock/shopwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopWriter
type ShopWriter interface {
	Create(shop entity.Shop) (entity.Shop, error)
	Update(shop entity.Shop) (entity.Shop, error)
	Delete(shop int32) error
}
