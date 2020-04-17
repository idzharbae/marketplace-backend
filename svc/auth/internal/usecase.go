package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

//go:generate mockgen -destination=usecase/ucmock/useruc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal UserUC
type UserUC interface {
	Get(user entity.User) (entity.User, error)
	GetWithPassword(user entity.User) (entity.User, error)
	List(req request.ListUser) ([]entity.User, error)
	GetShopsByProvince(province string) ([]entity.User, error)

	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(user entity.User) error
}

//go:generate mockgen -destination=usecase/ucmock/tokenuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal TokenUC
type TokenUC interface {
	Get(req entity.User) (entity.AuthToken, error)
	Refresh(req request.RefreshToken) (entity.AuthToken, error)
	Validate(req entity.AuthToken) error
}
