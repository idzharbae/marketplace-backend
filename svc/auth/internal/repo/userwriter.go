package repo

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
)

type UserWriter struct {
	db connection.Gormw
}

func NewUserWriter(db connection.Gormw) *UserWriter {
	return &UserWriter{db: db}
}

func (uw *UserWriter) Create(user entity.User) (entity.User, error) {
	var res model.User
	notFound := uw.db.Where("user_name=?", user.UserName).Or("email=?", user.Email).First(&res).RecordNotFound()
	if !notFound {
		return entity.User{}, errors.New("email or username already exists")
	}

	res = model.UserFromEntity(user)
	err := uw.db.Save(&res).Error()
	if err != nil {
		return entity.User{}, err
	}

	return res.ToEntity(), nil
}
func (uw *UserWriter) Update(user entity.User) (entity.User, error) {
	var res model.User
	notFound := uw.db.Where("id=?", user.ID).First(&res).RecordNotFound()
	if notFound {
		return entity.User{}, errors.New("user doesn't exists")
	}
	if res.Password != user.GetPasswordHash() {
		return entity.User{}, errors.New("wrong password")
	}

	res = model.User{
		ID:            res.ID,
		Name:          user.Name,
		UserName:      res.UserName,
		Email:         res.Email,
		Phone:         user.Phone,
		Type:          res.Type,
		PhotoURL:      res.PhotoURL,
		Province:      user.Address.Province,
		City:          user.Address.City,
		ZipCode:       user.Address.ZipCode,
		DetailAddress: user.Address.DetailAddress,
		Description:   user.Description,
		Password:      res.Password,
		CreatedAt:     res.CreatedAt,
		Saldo:         res.Saldo,
	}
	if user.NewPassword != "" {
		res.Password = user.GetNewPasswordHash()
	}
	if user.PhotoURL != "" {
		res.PhotoURL = user.PhotoURL
	}

	err := uw.db.Save(&res).Error()
	if err != nil {
		return entity.User{}, err
	}
	return res.ToEntity(), nil
}

func (uw *UserWriter) UpdateSaldo(req request.TopUp) (entity.User, error) {
	var user model.User
	findUserQuery := uw.db.Where("id=?", req.UserID).First(&user)
	if err := findUserQuery.Error(); err != nil {
		return entity.User{}, err
	}
	user.Saldo = user.Saldo + req.Amount
	err := uw.db.Save(&user).Error()
	if err != nil {
		return entity.User{}, err
	}
	return user.ToEntity(), nil
}

func (uw *UserWriter) TransferSaldo(req request.Transfer) (authproto.TransferSaldoResp, error) {
	var sender, receiver model.User
	err := uw.db.Where("id=?", req.SenderID).First(&sender).Error()
	if err != nil {
		return authproto.TransferSaldoResp{}, err
	}
	err = uw.db.Where("id=?", req.ReceiverID).First(&receiver).Error()
	if err != nil {
		return authproto.TransferSaldoResp{}, err
	}

	sender.Saldo -= req.TransferAmount
	receiver.Saldo += req.TransferAmount

	db := uw.db.Begin()
	err = db.Save(&sender).Error()
	if err != nil {
		return authproto.TransferSaldoResp{}, err
	}
	err = db.Save(&receiver).Error()
	if err != nil {
		db.Rollback()
		return authproto.TransferSaldoResp{}, err
	}
	db.Commit()
	return authproto.TransferSaldoResp{
		SenderId:      sender.ID,
		ReceiverId:    receiver.ID,
		SenderSaldo:   sender.Saldo,
		ReceiverSaldo: receiver.Saldo,
	}, nil
}

func (uw *UserWriter) DeleteByID(ID int64) error {
	return nil
}
