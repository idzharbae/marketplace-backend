package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/util"
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
	if user.UserName != "" {
		return u.userReader.GetByUserNameAndPassword(user)
	}
	if !util.IsEmail(user.Email) {
		return entity.User{}, errors.New("email is invalid")
	}
	return u.userReader.GetByEmailAndPassword(user)
}
func (u *User) List(req request.ListUser) ([]entity.User, error) {
	return nil, nil
}

func (u *User) Create(user entity.User) (entity.User, error) {
	if err := user.Validate(); err != nil {
		return entity.User{}, err
	}
	user.ID = 0
	return u.userWriter.Create(user)
}
func (u *User) Update(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (u *User) Delete(user entity.User) error {
	return nil
}
