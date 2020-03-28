package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
)

type UserWriter struct {
	db connection.Gormw
}

func NewUserWriter(db connection.Gormw) *UserWriter {
	return &UserWriter{db: db}
}

func (uw *UserWriter) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (uw *UserWriter) Update(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (uw *UserWriter) DeleteByID(ID int64) error {
	return nil
}
