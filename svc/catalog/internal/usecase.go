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

//go:generate mockgen -destination=usecase/ucmock/reviewuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/catalog/internal ReviewUC
type ReviewUC interface {
	Get(reveiwID int64) (entity.Review, error)
	List(req requests.ListReview) ([]entity.Review, error)

	Create(review entity.Review) (entity.Review, error)
	Update(review entity.Review) (entity.Review, error)
	Delete(review entity.Review) error
}
