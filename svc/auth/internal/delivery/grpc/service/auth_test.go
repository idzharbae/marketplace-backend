package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/usecase/ucmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testAuth struct {
	ctrl    *gomock.Controller
	tokenUC *ucmock.MockTokenUC
	userUC  *ucmock.MockUserUC
	unit    *AuthService
	ctx     context.Context
}

func newTestAuth() *testAuth {
	return &testAuth{}
}

func (ta *testAuth) Begin(t *testing.T) {
	ta.ctrl = gomock.NewController(t)
	ta.tokenUC = ucmock.NewMockTokenUC(ta.ctrl)
	ta.userUC = ucmock.NewMockUserUC(ta.ctrl)
	ta.unit = NewAuthService(ta.tokenUC, ta.userUC)
	ta.ctx = context.Background()
}

func (ta *testAuth) Finish() {
	ta.ctrl.Finish()
}

func TestAuthService_Login(t *testing.T) {
	test := newTestAuth()
	t.Run("user uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.LoginReq{
			UsernameOrEmail: "asdf",
			Password:        "asdaf",
		}

		test.userUC.EXPECT().GetWithPassword(gomock.Any()).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.Login(test.ctx, req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("error generating token, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.LoginReq{
			UsernameOrEmail: "asdf",
			Password:        "asdaf",
		}
		reqToken := entity.User{
			ID:       1337,
			UserName: "asdf",
			Email:    "asdf@asdf.com",
			Type:     1,
		}

		test.userUC.EXPECT().GetWithPassword(gomock.Any()).Return(entity.User{UserName: "asdf", Email: "asdf@asdf.com", ID: 1337, Type: 1}, nil)
		test.tokenUC.EXPECT().Get(reqToken).Return(entity.AuthToken{}, errors.New("error"))

		got, err := test.unit.Login(test.ctx, req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("success generation token, should return token", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.LoginReq{
			UsernameOrEmail: "asdf",
			Password:        "asdaf",
		}
		reqToken := entity.User{
			ID:       1337,
			UserName: "asdf",
			Email:    "asdf@asdf.com",
			Type:     1,
		}

		test.userUC.EXPECT().GetWithPassword(gomock.Any()).Return(entity.User{UserName: "asdf", Email: "asdf@asdf.com", ID: 1337, Type: 1}, nil)
		test.tokenUC.EXPECT().Get(reqToken).Return(entity.AuthToken{Token: "asdasdasd"}, nil)

		got, err := test.unit.Login(test.ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, "asdasdasd", got.Token)
	})
	t.Run("given email should fill email param", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.LoginReq{
			UsernameOrEmail: "asdf@asdf.com",
			Password:        "asdasdsad",
		}
		test.userUC.EXPECT().GetWithPassword(entity.User{
			Email:    req.UsernameOrEmail,
			Password: req.Password,
		}).Return(entity.User{}, errors.New("error"))
		test.unit.Login(context.Background(), req)
	})
	t.Run("given username should fill username param", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.LoginReq{
			UsernameOrEmail: "asdfghjk",
			Password:        "asdasdsad",
		}
		test.userUC.EXPECT().GetWithPassword(entity.User{
			UserName: req.UsernameOrEmail,
			Password: req.Password,
		}).Return(entity.User{}, errors.New("error"))
		test.unit.Login(context.Background(), req)
	})
}

func TestAuthService_Register(t *testing.T) {
	test := newTestAuth()
	t.Run("uc returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.RegisterReq{
			UserName: "asdf",
			Email:    "asdf",
			Phone:    "asdfg",
			Password: "asdasdasd",
			Type:     1,
			FullName: "asdasd",
		}

		test.userUC.EXPECT().Create(gomock.Any()).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.Register(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("uc returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := &authproto.RegisterReq{
			UserName: "asdf",
			Email:    "asdf",
			Phone:    "asdfg",
			Password: "asdasdasd",
			Type:     1,
			FullName: "asdasd",
		}

		test.userUC.EXPECT().Create(gomock.Any()).Return(entity.User{}, nil)

		got, err := test.unit.Register(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
