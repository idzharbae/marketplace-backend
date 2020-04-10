package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/converter"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userTest struct {
	ctrl   *gomock.Controller
	mockuc *ucmock.MockUserUC
	unit   *UserService
}

func newUserTest() *userTest {
	return &userTest{}
}

func (ut *userTest) Begin(t *testing.T) {
	ut.ctrl = gomock.NewController(t)
	ut.mockuc = ucmock.NewMockUserUC(ut.ctrl)
	ut.unit = NewUserService(ut.mockuc)
}

func (ut *userTest) Finish() {
	ut.ctrl.Finish()
}

func TestUserService_GetUser(t *testing.T) {
	test := newUserTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.GetUser(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})

	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.GetUserReq{
			Id:       123,
			UserName: "test",
			Email:    "test",
		}
		test.mockuc.EXPECT().Get(entity.User{
			ID: req.GetId(), UserName: req.GetUserName(),
			Email: req.GetEmail(),
		}).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.GetUser(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.GetUserReq{
			Id:       123,
			UserName: "test",
			Email:    "test",
		}
		test.mockuc.EXPECT().Get(entity.User{
			ID: req.GetId(), UserName: req.GetUserName(),
			Email: req.GetEmail(),
		}).Return(entity.User{
			ID:       123,
			Name:     "asdf",
			UserName: "test",
			Email:    "test",
			Phone:    "Asdasd",
			PhotoURL: "asdasdasd",
			Type:     1,
		}, nil)

		got, err := test.unit.GetUser(context.Background(), req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, req.GetId(), got.GetId())
		assert.Equal(t, req.GetUserName(), got.GetUserName())
		assert.Equal(t, req.GetEmail(), got.GetEmail())
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	test := newUserTest()
	t.Run("given nil param should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		got, err := test.unit.UpdateUser(context.Background(), nil)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.User{
			Id:       1,
			Name:     "asd",
			UserName: "asdf",
			Email:    "Asdf",
			Phone:    "Asdasd",
			PhotoUrl: "Asdasd",
			Type:     1,
			Password: "asdasd",
		}
		test.mockuc.EXPECT().Update(entity.User{
			ID:       req.GetId(),
			Name:     req.GetName(),
			UserName: req.GetUserName(),
			Email:    req.GetEmail(),
			Phone:    req.GetPhone(),
			PhotoURL: req.GetPhotoUrl(),
			Password: req.GetPassword(),
			Type:     req.GetType(),
		}).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.UpdateUser(context.Background(), req)
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
	t.Run("no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.User{
			Id:       1,
			Name:     "asd",
			UserName: "asdf",
			Email:    "Asdf",
			Phone:    "Asdasd",
			PhotoUrl: "Asdasd",
			Type:     1,
			Password: "asdasd",
		}
		test.mockuc.EXPECT().Update(entity.User{
			ID:       req.GetId(),
			Name:     req.GetName(),
			UserName: req.GetUserName(),
			Email:    req.GetEmail(),
			Phone:    req.GetPhone(),
			PhotoURL: req.GetPhotoUrl(),
			Password: req.GetPassword(),
			Type:     req.GetType(),
		}).Return(converter.UserProtoToEntity(req), nil)

		got, err := test.unit.UpdateUser(context.Background(), req)
		assert.NotNil(t, got)
		assert.Nil(t, err)
	})
}
