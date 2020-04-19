package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
)

//go:generate mockgen -destination=gateway/gatewaymock/cataloggateway_mock.go -package=gatewaymock github.com/idzharbae/marketplace-backend/svc/transaction/internal CatalogGateway
type CatalogGateway interface {
	GetProductByID(productID int64) (entity.Product, error)
}
