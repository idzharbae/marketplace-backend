package model

import "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"

type User struct {
	ID       int64
	Name     string
	UserName string
	Email    string
	Phone    string
	Password string
	Type     int32
}

func UserFromEntity(user entity.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.GetPasswordHash(),
		Type:     user.Type,
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
	}
}
