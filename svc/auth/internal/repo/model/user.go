package model

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"time"
)

type User struct {
	ID            int64
	Name          string
	UserName      string
	Email         string
	Phone         string
	Password      string
	Type          int32
	PhotoURL      string
	Province      string
	City          string
	ZipCode       int32
	DetailAddress string
	Description   string
	Saldo         int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func UserFromEntity(user entity.User) User {
	return User{
		ID:            user.ID,
		Name:          user.Name,
		UserName:      user.UserName,
		Email:         user.Email,
		Phone:         user.Phone,
		Password:      user.GetPasswordHash(),
		Type:          user.Type,
		PhotoURL:      user.PhotoURL,
		Province:      user.Address.Province,
		City:          user.Address.City,
		ZipCode:       user.Address.ZipCode,
		DetailAddress: user.Address.DetailAddress,
		Description:   user.Description,
		Saldo:         user.Saldo,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

func (u User) TableName() string {
	return "user_account"
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.UserName,
		Email:    u.Email,
		Phone:    u.Phone,
		Type:     u.Type,
		PhotoURL: u.PhotoURL,
		Address: entity.Address{
			Province:      u.Province,
			City:          u.City,
			DetailAddress: u.DetailAddress,
			ZipCode:       u.ZipCode,
		},
		Description: u.Description,
		Saldo:       u.Saldo,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
