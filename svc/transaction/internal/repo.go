package internal

import "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"

//go:generate mockgen -destination=repo/repomock/cartreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal CartReader
type CartReader interface {
	ListByUserID(userID int64) ([]entity.Cart, error)
}

//go:generate mockgen -destination=repo/repomock/cartwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/transaction/internal CartWriter
type CartWriter interface {
	Create(cart entity.Cart) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	DeleteByID(cartID int64) error
}
