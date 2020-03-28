package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
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
