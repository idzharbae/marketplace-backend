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
	GetTotalByShopID(shopID int32) (int32, error)
}

//go:generate mockgen -destination=repo/repomock/productwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ProductWriter
type ProductWriter interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	DeleteByID(productID int32) error
	DeleteBySlug(productSlug string) error
}

//go:generate mockgen -destination=repo/repomock/reviewreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ReviewReader
type ReviewReader interface {
	ListByProductID(productID int64, pagination requests.Pagination) ([]entity.Review, error)
	ListByShopID(shopID int64, pagination requests.Pagination) ([]entity.Review, error)

	GetByID(reviewID int64) (entity.Review, error)
	GetTotalAndAverageByShopID(shopID int64) (requests.TotalAndAverageReview, error)
	GetTotalAndAverageByProductID(productID int64) (requests.TotalAndAverageReview, error)
}

//go:generate mockgen -destination=repo/repomock/reviewwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/catalog/internal ReviewWriter
type ReviewWriter interface {
	Create(review entity.Review) (entity.Review, error)
	Update(review entity.Review) (entity.Review, error)
	Delete(review entity.Review) error
}
