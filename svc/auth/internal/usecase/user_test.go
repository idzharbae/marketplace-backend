package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/repomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userTest struct {
	ctrl   *gomock.Controller
	reader *repomock.MockUserReader
	writer *repomock.MockUserWriter
	unit   internal.UserUC
}

func newUserTest() *userTest {
	return &userTest{}
}

func (ut *userTest) Begin(t *testing.T) {
	ut.ctrl = gomock.NewController(t)
	ut.reader = repomock.NewMockUserReader(ut.ctrl)
	ut.writer = repomock.NewMockUserWriter(ut.ctrl)
	ut.unit = NewUser(ut.reader, ut.writer)
}

func (ut *userTest) Finish() {
	ut.ctrl.Finish()
}

func TestUser_Get(t *testing.T) {
	test := newUserTest()
	t.Run("given username, repo error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			UserName: "username",
			Password: "asdfghj",
		}

		test.reader.EXPECT().GetByUserNameAndPassword(req).Return(entity.User{}, errors.New("error"))
		got, err := test.unit.Get(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("given email, email invalid, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			Email:    "username",
			Password: "asdfghj",
		}

		got, err := test.unit.Get(req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
}

func TestUser_Create(t *testing.T) {
	test := newUserTest()
	t.Run("given invalid email, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdasfasfas",
			UserName: "asdasdasdas",
			Email:    "asadsasdadf",
			Phone:    "123123123",
			Password: "asdasdasdsd",
			Type:     1,
		}
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("given invalid phone number, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdasfasfas",
			UserName: "asdasdasdas",
			Email:    "idzharbae@gmail.com",
			Phone:    "12312312a3",
			Password: "asdasdasdsd",
			Type:     1,
		}
		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       0,
			Name:     "asdasfasfas",
			UserName: "asdasdasdas",
			Email:    "idzharbae@gmail.com",
			Phone:    "123123122133",
			Password: "asdasdasdsd",
			Type:     1,
		}

		test.writer.EXPECT().Create(gomock.Any()).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.Create(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
}
