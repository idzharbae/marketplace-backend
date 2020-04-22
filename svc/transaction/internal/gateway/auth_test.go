package gateway

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/gateway/connection/connectionmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type authTest struct {
	ctrl   *gomock.Controller
	client *connectionmock.MockAuth
	unit   internal.AuthGateway
}

func newAuthTest() *authTest {
	return &authTest{}
}

func (at *authTest) Begin(t *testing.T) {
	at.ctrl = gomock.NewController(t)
	at.client = connectionmock.NewMockAuth(at.ctrl)
	at.unit = NewAuth(at.client)
}

func (at *authTest) Finish() {
	at.ctrl.Finish()
}

func TestAuth_GetUserByID(t *testing.T) {
	test := newAuthTest()
	t.Run("client return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.client.EXPECT().GetUser(context.Background(), &authproto.GetUserReq{Id: req}).Return(nil, errors.New("error"))
		got, err := test.unit.GetUserByID(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("client return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(123)
		test.client.EXPECT().GetUser(context.Background(), &authproto.GetUserReq{Id: req}).Return(&authproto.User{Id: 123, Saldo: 1337}, nil)
		got, err := test.unit.GetUserByID(req)
		assert.Nil(t, err)
		assert.Equal(t, entity.User{ID: 123, Saldo: 1337}, got)
	})
}

func TestAuth_UpdateUserSaldo(t *testing.T) {
	test := newAuthTest()
	t.Run("client returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		test.client.EXPECT().UpdateSaldo(context.Background(), &authproto.TopUpReq{Amount: 456, UserId: 123}).Return(nil, errors.New("error"))
		got, err := test.unit.UpdateUserSaldo(123, 456)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("client returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		test.client.EXPECT().UpdateSaldo(context.Background(), &authproto.TopUpReq{Amount: 456, UserId: 123}).Return(&authproto.TopUpResp{
			UserId: 123,
			Saldo:  1456,
		}, nil)
		got, err := test.unit.UpdateUserSaldo(123, 456)
		assert.Nil(t, err)
		assert.Equal(t, entity.User{ID: 123, Saldo: 1456}, got)
	})
}
