package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
)

func UserEntityToProto(user entity.User) *authproto.User {
	return &authproto.User{
		Id:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
		PhotoUrl: user.PhotoURL,
		Type:     user.Type,
	}
}

func UserProtoToEntity(in *authproto.User) entity.User {
	return entity.User{
		ID:       in.GetId(),
		Name:     in.GetName(),
		UserName: in.GetUserName(),
		Email:    in.GetEmail(),
		Phone:    in.GetPhone(),
		PhotoURL: in.GetPhotoUrl(),
		Password: in.GetPassword(),
		Type:     in.GetType(),
	}
}

func RegisterReqToEntity(in *authproto.RegisterReq) entity.User {
	return entity.User{
		Name:     in.GetFullName(),
		UserName: in.GetUserName(),
		Email:    in.GetEmail(),
		Phone:    in.GetPhone(),
		Password: in.GetPassword(),
		Type:     in.GetType(),
		PhotoURL: in.GetPhotoUrl(),
	}
}
