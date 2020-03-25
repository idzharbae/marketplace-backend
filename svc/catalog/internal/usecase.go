package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

//go:generate mockgen -destination=usecase/ucmock/productuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductUC
type ProductUC interface {
	List(req requests.ListProduct) ([]entity.Product, error)
	Get(product entity.Product) (entity.Product, error)

	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(product entity.Product) error
}

//go:generate mockgen -destination=usecase/ucmock/shopuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/catalog/internal ShopUC
type ShopUC interface {
	List(req requests.ListShop) ([]entity.Shop, error)
	Get(shop entity.Shop) (entity.Shop, error)

	Create(shop entity.Shop) (entity.Shop, error)
	Update(shop entity.Shop) (entity.Shop, error)
	Delete(shop entity.Shop) error
}
