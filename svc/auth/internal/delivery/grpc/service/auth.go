package service

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/util"
)

type AuthService struct {
	tokenUC internal.TokenUC
	userUC  internal.UserUC
}

func NewAuthService(tokenUC internal.TokenUC, userUC internal.UserUC) *AuthService {
	return &AuthService{tokenUC: tokenUC, userUC: userUC}
}

func (as *AuthService) Login(ctx context.Context, in *authproto.LoginReq) (*authproto.LoginResp, error) {
	req := entity.User{
		Password: in.GetPassword(),
	}
	if util.IsEmail(in.GetUsernameOrEmail()) {
		req.Email = in.GetUsernameOrEmail()
	} else {
		req.UserName = in.GetUsernameOrEmail()
	}
	user, err := as.userUC.GetWithPassword(req)
	if err != nil {
		return nil, err
	}
	token, err := as.tokenUC.Get(user)
	if err != nil {
		return nil, err
	}

	return &authproto.LoginResp{
		Token: token.Token,
	}, nil
}

func (as *AuthService) Register(ctx context.Context, in *authproto.RegisterReq) (*authproto.User, error) {
	res, err := as.userUC.Create(converter.RegisterReqToEntity(in))
	if err != nil {
		return nil, err
	}
	return converter.UserEntityToProto(res), nil
}

func (as *AuthService) RefreshToken(ctx context.Context, in *authproto.RefreshTokenReq) (*authproto.RefreshTokenResp, error) {
	token, err := as.tokenUC.Refresh(request.RefreshToken{CurrentToken: in.GetToken()})
	if err != nil {
		return nil, err
	}
	return &authproto.RefreshTokenResp{
		Token: token.Token,
	}, nil
}
