package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type User struct {
	userReader internal.UserReader
	userWriter internal.UserWriter
}

func NewUser(userReader internal.UserReader, userWriter internal.UserWriter) *User {
	return &User{
		userReader: userReader,
		userWriter: userWriter,
	}
}

func (u *User) Get(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (u *User) List(req request.ListUser) ([]entity.User, error) {
	return nil, nil
}

func (u *User) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (u *User) Update(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (u *User) Delete(user entity.User) error {
	return nil
}
