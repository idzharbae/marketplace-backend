package internal

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

//go:generate mockgen -destination=repo/repomock/userreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal UserReader
type UserReader interface {
	ListAll(req request.ListUser) ([]entity.User, error)

	GetByID(ID int64) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByUserName(username string) (entity.User, error)
	GetShopsByProvince(province string) ([]entity.User, error)

	GetByEmailAndPassword(user entity.User) (entity.User, error)
	GetByUserNameAndPassword(user entity.User) (entity.User, error)
}

//go:generate mockgen -destination=repo/repomock/userwriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal UserWriter
type UserWriter interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	UpdateSaldo(req request.TopUp) (entity.User, error)
	TransferSaldo(req request.Transfer) (authproto.TransferSaldoResp, error)
	DeleteByID(ID int64) error
}

//go:generate mockgen -destination=repo/repomock/saldohistoryreader_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal SaldoHistoryReader
type SaldoHistoryReader interface {
	ListByUserID(userID int64) ([]entity.SaldoHistory, error)
}

//go:generate mockgen -destination=repo/repomock/saldohistorywriter_mock.go -package=repomock github.com/idzharbae/marketplace-backend/svc/auth/internal SaldoHistoryWriter
type SaldoHistoryWriter interface {
	Create(history entity.SaldoHistory) (entity.SaldoHistory, error)
}
