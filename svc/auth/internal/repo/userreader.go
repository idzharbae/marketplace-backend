package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type UserReader struct {
	db connection.Gormw
}

func NewUserReader(db connection.Gormw) *UserReader {
	return &UserReader{db: db}
}

func (ur *UserReader) ListAll(req request.ListUser) ([]entity.User, error) {
	return nil, nil
}
func (ur *UserReader) GetByID(ID int64) (entity.User, error) {
	return entity.User{}, nil
}
func (ur *UserReader) GetByUserNameAndPassword(req entity.User) (entity.User, error) {
	var user model.User
	err := ur.db.Where("user_name=?", req.UserName).Where("password=?", req.GetPasswordHash()).First(&user).Error()
	if err != nil {
		return entity.User{}, err
	}
	return user.ToEntity(), nil
}
func (ur *UserReader) GetByEmailAndPassword(req entity.User) (entity.User, error) {
	var user model.User
	err := ur.db.Where("email=?", req.Email).Where("password=?", req.GetPasswordHash()).First(&user).Error()
	if err != nil {
		return entity.User{}, err
	}
	return user.ToEntity(), nil
}
