package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

//go:generate mockgen -destination=usecase/ucmock/productuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductUC
type ProductUC interface {
	List(req requests.ListProduct) ([]entity.Product, error)
	GetByID(productID int32) (entity.Product, error)
	GetBySlug(slug string) (entity.Product, error)
	GetByShopID(shopID int32) ([]entity.Product, error)

	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(productID int32) error
}

//go:generate mockgen -destination=usecase/ucmock/shopuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopUC
type ShopUC interface {
	List(req requests.ListShop) ([]entity.Shop, error)
	GetByID(shopID int32) (entity.Shop, error)
	GetBySlug(slug string) (entity.Shop, error)

	Create(shop entity.Shop) (entity.Shop, error)
	Update(shop entity.Shop) (entity.Shop, error)
	Delete(shop int32) error
}
