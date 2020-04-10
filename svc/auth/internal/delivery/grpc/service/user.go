package service

import (
	"context"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"

	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
)

type UserService struct {
	UserUC internal.UserUC
}

func NewUserService(userUC internal.UserUC) *UserService {
	return &UserService{UserUC: userUC}
}

func (us *UserService) GetUser(ctx context.Context, in *authproto.GetUserReq) (*authproto.User, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := us.UserUC.Get(entity.User{
		ID:       in.GetId(),
		UserName: in.GetUserName(),
		Email:    in.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	return converter.UserEntityToProto(res), nil
}

func (us *UserService) UpdateUser(ctx context.Context, in *authproto.User) (*authproto.User, error) {
	if in == nil {
		return nil, errors.New("parameter should not be nil")
	}
	res, err := us.UserUC.Update(converter.UserProtoToEntity(in))
	if err != nil {
		return nil, err
	}
	return converter.UserEntityToProto(res), nil
}
