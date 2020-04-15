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
