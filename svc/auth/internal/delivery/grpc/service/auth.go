package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
)

type AuthService struct {
	tokenUC internal.TokenUC
}

func NewAuthService(tokenUC internal.TokenUC) *AuthService {
	return &AuthService{tokenUC: tokenUC}
}

func (as *AuthService) Login(ctx context.Context, in *authproto.LoginReq) (*authproto.LoginResp, error) {
	return nil, nil
}

func (as *AuthService) Register(ctx context.Context, in *authproto.RegisterReq) (*authproto.RegisterResp, error) {
	return nil, nil
}

func (as *AuthService) RefreshToken(ctx context.Context, in *authproto.RefreshTokenReq) (*authproto.RefreshTokenResp, error) {
	return nil, nil
}
