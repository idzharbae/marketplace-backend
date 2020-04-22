package gateway

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/converter"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection"
)

type Auth struct {
	conn connection.Auth
}

func NewAuth(conn connection.Auth) *Auth {
	return &Auth{conn: conn}
}

func (a *Auth) GetUserByID(userID int64) (entity.User, error) {
	user, err := a.conn.GetUser(context.Background(), &authproto.GetUserReq{Id: userID})
	if err != nil {
		return entity.User{}, err
	}
	return converter.UserProtoToEntity(user), nil
}
func (a *Auth) UpdateUserSaldo(userID int64, changeAmount int64) (entity.User, error) {
	res, err := a.conn.UpdateSaldo(context.Background(), &authproto.TopUpReq{UserId: userID, Amount: changeAmount})
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{ID: res.GetUserId(), Saldo: res.GetSaldo()}, nil
}
