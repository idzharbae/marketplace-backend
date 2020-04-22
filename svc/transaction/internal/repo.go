package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
)

//go:generate mockgen -destination=repo/repomock/cartreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal CartReader
type CartReader interface {
	ListByUserID(userID int64) ([]entity.Cart, error)
	GetByIDs(cartID ...int64) ([]entity.Cart, error)
}

//go:generate mockgen -destination=repo/repomock/cartwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal CartWriter
type CartWriter interface {
	Create(cart entity.Cart) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	DeleteByID(cartID int64) error
}

//go:generate mockgen -destination=repo/repomock/orderreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal OrderReader
type OrderReader interface {
	ListByUserID(userID int64) ([]entity.Order, error)
	ListByShopID(shopID int64) ([]entity.Order, error)
	GetByID(orderID int64) (entity.Order, error)
}

//go:generate mockgen -destination=repo/repomock/orderwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal OrderWriter
type OrderWriter interface {
	CreateFromCartsAndSubstractCustomerSaldo(req request.CreateOrderReq) ([]entity.Order, error)
	UpdateOrderStatusToOnShipment(orderID, shopID int64) (entity.Order, error)
	UpdateOrderStatusAndAddShopSaldo(order entity.Order) (entity.Order, error)
}
