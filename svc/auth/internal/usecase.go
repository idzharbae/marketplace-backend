package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

//go:generate mockgen -destination=usecase/ucmock/useruc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal UserUC
type UserUC interface {
	Get(user entity.User) (entity.User, error)
	GetWithPassword(user entity.User) (entity.User, error)
	List(req request.ListUser) ([]entity.User, error)
	GetShopsByProvince(province string) ([]entity.User, error)

	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(user entity.User) error
}

//go:generate mockgen -destination=usecase/ucmock/tokenuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal TokenUC
type TokenUC interface {
	Get(req entity.User) (entity.AuthToken, error)
	Refresh(req request.RefreshToken) (entity.AuthToken, error)
	Validate(req entity.AuthToken) error
}

//go:generate mockgen -destination=usecase/ucmock/paymentuc_mock.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal PaymentUC
type PaymentUC interface {
	TopUp(req request.TopUp) (entity.User, error)
	UpdateSaldo(req request.TopUp) (entity.User, error)
	Transfer(req request.Transfer) (authproto.TransferSaldoResp, error)
}

//go:generate mockgen -destination=usecase/ucmock/saldohistory_uc.go -package=ucmock github.com/idzharbae/marketplace-backend/svc/auth/internal SaldoHistoryUC
type SaldoHistoryUC interface {
	List(userID int64) ([]entity.SaldoHistory, error)
	Create(req entity.SaldoHistory) (entity.SaldoHistory, error)
}
