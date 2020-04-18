package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userWriterTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.UserWriter
}

func newUserWriterTest() *userWriterTest {
	return &userWriterTest{}
}

func (uw *userWriterTest) Begin(t *testing.T) {
	uw.ctrl = gomock.NewController(t)
	uw.db = gormmock.NewMockGormw(uw.ctrl)
	uw.unit = NewUserWriter(uw.db)
}

func (uw *userWriterTest) Finish() {
	uw.ctrl.Finish()
}

func TestUserWriter_Create(t *testing.T) {
	test := newUserWriterTest()
	t.Run("username already exists should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdf",
			UserName: "asdf",
			Email:    "asdf",
			Phone:    "Asdf",
			Password: "asd",
			Type:     1,
		}

		test.db.EXPECT().Where("user_name=?", req.UserName).Return(test.db)
		test.db.EXPECT().Or("email=?", req.UserName).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.User{
				ID:       1337,
				Name:     "asdasda",
				UserName: "asdsada",
			}
			return test.db
		}).Return(test.db)
		test.db.EXPECT().RecordNotFound().Return(false)

		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db returns error when saving, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdf",
			UserName: "asdf",
			Email:    "asdf",
			Phone:    "Asdf",
			Password: "asd",
			Type:     1,
		}

		test.db.EXPECT().Where("user_name=?", req.UserName).Return(test.db)
		test.db.EXPECT().Or("email=?", req.UserName).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.User{
				ID:       1337,
				Name:     "asdasda",
				UserName: "asdsada",
			}
			return test.db
		}).Return(test.db)
		test.db.EXPECT().RecordNotFound().Return(true)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db returns no error when saving, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdf",
			UserName: "asdf",
			Email:    "asdf",
			Phone:    "Asdf",
			Password: "asd",
			Type:     1,
		}

		test.db.EXPECT().Where("user_name=?", req.UserName).Return(test.db)
		test.db.EXPECT().Or("email=?", req.UserName).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.User{
				ID:       1337,
				Name:     "asdasda",
				UserName: "asdsada",
			}
			return test.db
		}).Return(test.db)
		test.db.EXPECT().RecordNotFound().Return(true)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.User{
				ID:       1,
				Name:     user.Name,
				UserName: user.UserName,
			}
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Create(req)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), got.ID)
	})
}

func TestUserWriter_Update(t *testing.T) {
	test := newUserWriterTest()
	t.Run("record not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "Asdf",
			UserName: "asdff",
			Email:    "asdasd",
			Phone:    "Asdasd",
			PhotoURL: "Asdasd",
			Password: "asdasd",
			Type:     1,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = model.User{
				ID:            req.ID,
				Name:          "asdf",
				UserName:      "asdf",
				Email:         "asdf",
				Phone:         "asdf",
				Password:      req.GetPasswordHash(),
				Type:          1,
				PhotoURL:      "asdf",
				Province:      "asdf",
				City:          "asdf",
				ZipCode:       123123,
				DetailAddress: "asda",
				Description:   "asdsad",
			}
			return test.db
		})
		test.db.EXPECT().RecordNotFound().Return(true)

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("wrong password, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "Asdf",
			UserName: "asdff",
			Email:    "asdasd",
			Phone:    "Asdasd",
			PhotoURL: "Asdasd",
			Password: "asdasd",
			Type:     1,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = model.User{
				ID:            req.ID,
				Name:          "asdf",
				UserName:      "asdf",
				Email:         "asdf",
				Phone:         "asdf",
				Password:      "asdasfasf",
				Type:          1,
				PhotoURL:      "asdf",
				Province:      "asdf",
				City:          "asdf",
				ZipCode:       123123,
				DetailAddress: "asda",
				Description:   "asdsad",
			}
			return test.db
		})
		test.db.EXPECT().RecordNotFound().Return(false)

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "Asdf",
			UserName: "asdff",
			Email:    "asdasd",
			Phone:    "Asdasd",
			PhotoURL: "Asdasd",
			Password: "asdasd",
			Type:     1,
		}

		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = model.User{
				ID:            req.ID,
				Name:          "asdf",
				UserName:      "asdf",
				Email:         "asdf",
				Phone:         "asdf",
				Password:      req.GetPasswordHash(),
				Type:          1,
				PhotoURL:      "asdf",
				Province:      "asdf",
				City:          "asdf",
				ZipCode:       123123,
				DetailAddress: "asda",
				Description:   "asdsad",
			}
			return test.db
		})
		test.db.EXPECT().RecordNotFound().Return(false)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("repo returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "Asdf",
			UserName: "asdff",
			Email:    "asdasd",
			Phone:    "Asdasd",
			PhotoURL: "Asdasd",
			Password: "asdasd",
			Type:     1,
		}
		test.db.EXPECT().Where("id=?", req.ID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = model.User{
				ID:            req.ID,
				Name:          "asdf",
				UserName:      "asdf",
				Email:         "asdf",
				Phone:         "asdf",
				Password:      req.GetPasswordHash(),
				Type:          1,
				PhotoURL:      "asdf",
				Province:      "asdf",
				City:          "asdf",
				ZipCode:       123123,
				DetailAddress: "asda",
				Description:   "asdsad",
			}
			return test.db
		})
		test.db.EXPECT().RecordNotFound().Return(false)
		test.db.EXPECT().Save(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = model.UserFromEntity(req)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.Update(req)
		assert.Nil(t, err)
		assert.Equal(t, req.Name, got.Name)
	})
}

func TestUserWriter_UpdateSaldo(t *testing.T) {
	test := newUserWriterTest()
	t.Run("db return error when searching user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.TopUp{
			UserID: 123,
			Amount: 1337,
		}

		test.db.EXPECT().Where("id=?", req.UserID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.UpdateSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db return error when saving user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.TopUp{
			UserID: 123,
			Amount: 1337,
		}
		userModel := model.User{
			ID:    req.UserID,
			Saldo: 1337,
		}

		test.db.EXPECT().Where("id=?", req.UserID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = userModel
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("asdfg"))

		got, err := test.unit.UpdateSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db return error when saving user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.TopUp{
			UserID: 123,
			Amount: 1337,
		}
		userModel := model.User{
			ID:    req.UserID,
			Saldo: 1337,
		}

		test.db.EXPECT().Where("id=?", req.UserID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = userModel
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.UpdateSaldo(req)
		assert.Nil(t, err)
		assert.Equal(t, req.Amount+userModel.Saldo, got.Saldo)
	})
}

func TestUserWriter_TransferSaldo(t *testing.T) {
	test := newUserWriterTest()
	t.Run("db return error when searching sender user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       12,
			ReceiverID:     15,
			TransferAmount: 1234,
		}

		test.db.EXPECT().Where("id=?", req.SenderID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.TransferSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, authproto.TransferSaldoResp{}, got)
	})
	t.Run("db return error when searching receiver user, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       12,
			ReceiverID:     15,
			TransferAmount: 1234,
		}
		sender := model.User{ID: 12, Saldo: 1339}

		test.db.EXPECT().Where("id=?", req.SenderID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = sender
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("id=?", req.ReceiverID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.TransferSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, authproto.TransferSaldoResp{}, got)
	})
	t.Run("db return error when saving sender new saldo, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       12,
			ReceiverID:     15,
			TransferAmount: 1234,
		}
		sender := model.User{ID: 12, Saldo: 1339}
		receiver := model.User{ID: 12, Saldo: 4567}

		test.db.EXPECT().Where("id=?", req.SenderID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = sender
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("id=?", req.ReceiverID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = receiver
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.TransferSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, authproto.TransferSaldoResp{}, got)
	})
	t.Run("db return error when saving receiver new saldo, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       12,
			ReceiverID:     15,
			TransferAmount: 1234,
		}
		sender := model.User{ID: 12, Saldo: 1339}
		receiver := model.User{ID: 12, Saldo: 4567}

		test.db.EXPECT().Where("id=?", req.SenderID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = sender
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("id=?", req.ReceiverID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = receiver
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))
		test.db.EXPECT().Rollback()

		got, err := test.unit.TransferSaldo(req)
		assert.NotNil(t, err)
		assert.Equal(t, authproto.TransferSaldoResp{}, got)
	})
	t.Run("transfer success, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := request.Transfer{
			SenderID:       12,
			ReceiverID:     15,
			TransferAmount: 1234,
		}
		sender := model.User{ID: 12, Saldo: 1339}
		receiver := model.User{ID: 12, Saldo: 4567}

		test.db.EXPECT().Where("id=?", req.SenderID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = sender
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Where("id=?", req.ReceiverID).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = receiver
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Begin().Return(test.db)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Save(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(nil)
		test.db.EXPECT().Commit()

		got, err := test.unit.TransferSaldo(req)
		assert.Nil(t, err)
		assert.Equal(t, sender.Saldo-req.TransferAmount, got.SenderSaldo)
		assert.Equal(t, receiver.Saldo+req.TransferAmount, got.ReceiverSaldo)
	})
}
