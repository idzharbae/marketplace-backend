package entity

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	ID       int64
	Name     string
	UserName string
	Email    string
	Phone    string
	Password string
	Type     int32
}

func (u User) GetPasswordHash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))
}
