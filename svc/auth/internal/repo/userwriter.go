package repo

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
)

type UserWriter struct {
	db connection.Gormw
}

func NewUserWriter(db connection.Gormw) *UserWriter {
	return &UserWriter{db: db}
}

func (uw *UserWriter) Create(user entity.User) (entity.User, error) {
	var res model.User
	notFound := uw.db.Where("user_name=?", user.UserName).Or("email=?", user.Email).First(&res).RecordNotFound()
	if !notFound {
		return entity.User{}, errors.New("email or username already exists")
	}

	res = model.UserFromEntity(user)
	err := uw.db.Save(&res).Error()
	if err != nil {
		return entity.User{}, err
	}

	return res.ToEntity(), nil
}
func (uw *UserWriter) Update(user entity.User) (entity.User, error) {
	var res model.User
	notFound := uw.db.Where("id=?", user.ID).First(&res).RecordNotFound()
	if notFound {
		return entity.User{}, errors.New("user doesn't exists")
	}
	if res.Password != user.GetPasswordHash() {
		return entity.User{}, errors.New("wrong password")
	}

	res = model.User{
		ID:            res.ID,
		Name:          user.Name,
		UserName:      res.UserName,
		Email:         res.Email,
		Phone:         user.Phone,
		Type:          res.Type,
		PhotoURL:      user.PhotoURL,
		Province:      user.Address.Province,
		City:          user.Address.City,
		ZipCode:       user.Address.ZipCode,
		DetailAddress: user.Address.DetailAddress,
		Description:   user.Description,
		Password:      res.Password,
		CreatedAt:     res.CreatedAt,
	}
	if user.NewPassword != "" {
		res.Password = user.GetNewPasswordHash()
	}

	err := uw.db.Save(&res).Error()
	if err != nil {
		return entity.User{}, err
	}
	return res.ToEntity(), nil
}
func (uw *UserWriter) DeleteByID(ID int64) error {
	return nil
}
