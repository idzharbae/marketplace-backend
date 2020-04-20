package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
)

func UserProtoToEntity(user *authproto.User) entity.User {
	return entity.User{
		ID:    user.GetId(),
		Saldo: user.GetSaldo(),
	}
}
