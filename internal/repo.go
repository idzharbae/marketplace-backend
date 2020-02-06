package internal

import (
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/requests"
)

type ProductReader interface {
	ListProducts(req requests.ListProduct) ([]entity.Product, error)
}

type ProductWriter interface {

}