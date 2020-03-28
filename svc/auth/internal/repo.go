package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

//go:generate mockgen -destination=repo/repomock/userreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal UserReader
type UserReader interface {
	ListAll(req request.ListUser) ([]entity.User, error)
	GetByID(ID int64) (entity.User, error)
}

//go:generate mockgen -destination=repo/repomock/userwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal UserWriter
type UserWriter interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	DeleteByID(ID int64) error
}