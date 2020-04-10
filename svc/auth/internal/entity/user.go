package entity

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/util"
)

type User struct {
	ID       int64
	Name     string
	UserName string
	Email    string
	Phone    string
	PhotoURL string
	Password string
	Type     int32
}

func (u User) GetPasswordHash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))
}

func (u User) Validate() error {
	if !util.IsEmail(u.Email) {
		return errors.New("invalid email")
	}
	if !util.IsNumeric(u.Phone) {
		return errors.New("invalid phone number")
	}
	if len(u.Password) < 6 {
		return errors.New("password too short")
	}
	if len(u.UserName) < 6 {
		return errors.New("username too short")
	}
	if len(u.Email) > 60 {
		return errors.New("email too long")
	}
	if len(u.Password) > 60 {
		return errors.New("password too long")
	}
	if len(u.UserName) > 60 {
		return errors.New("username too long")
	}
	return nil
}
