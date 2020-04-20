package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
)

//go:generate mockgen -destination=gateway/gatewaymock/cataloggateway_mock.go -package=gatewaymock github.com/idzharbae/marketplace-backend/svc/transaction/internal CatalogGateway
type CatalogGateway interface {
	GetProductByID(productID int64) (entity.Product, error)
	GetProductsByID(productIDs []int64) ([]entity.Product, error)
}

//go:generate mockgen -destination=gateway/gatewaymock/authgateway_mock.go -package=gatewaymock github.com/idzharbae/marketplace-backend/svc/transaction/internal AuthGateway
type AuthGateway interface {
	GetUserByID(userID int64) (entity.User, error)
	UpdateUserSaldo(userID int64, changeAmount int64) (entity.User, error)
}
