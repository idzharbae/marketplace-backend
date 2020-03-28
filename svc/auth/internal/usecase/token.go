package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type Token struct {
	userReader internal.UserReader
}

func NewToken(userReader internal.UserReader) *Token {
	return &Token{
		userReader: userReader,
	}
}

func (t *Token) Get(req request.GetToken) (entity.AuthToken, error) {
	return entity.AuthToken{}, nil
}

func (t *Token) Refresh(req request.RefreshToken) (entity.AuthToken, error) {
	return entity.AuthToken{}, nil
}
