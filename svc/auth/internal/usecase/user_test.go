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

func TestUser_GetWithPassword(t *testing.T) {
	test := newUserTest()
	t.Run("given username, repo error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			UserName: "username",
			Password: "asdfghj",
		}

		test.reader.EXPECT().GetByUserNameAndPassword(req).Return(entity.User{}, errors.New("error"))
		got, err := test.unit.GetWithPassword(req)
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

		got, err := test.unit.GetWithPassword(req)

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

func TestUser_Get(t *testing.T) {
	test := newUserTest()
	t.Run("given userID should call GetByUserID", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{ID: 1337}
		test.reader.EXPECT().GetByID(req.ID).Return(entity.User{ID: 1337, Name: "asdfg,"}, errors.New("error"))

		got, err := test.unit.Get(req)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		assert.Equal(t, "asdfg,", got.Name)
	})
	t.Run("given email should call GetByEmail", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{Email: "asdf"}
		test.reader.EXPECT().GetByEmail(req.Email).Return(entity.User{ID: 1337, Name: "asdfg,"}, errors.New("error"))

		got, err := test.unit.Get(req)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		assert.Equal(t, "asdfg,", got.Name)
	})
	t.Run("given username should call GetByUserName", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{UserName: "asdasd"}
		test.reader.EXPECT().GetByUserName(req.UserName).Return(entity.User{ID: 1337, Name: "asdfg,"}, errors.New("error"))

		got, err := test.unit.Get(req)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		assert.Equal(t, "asdfg,", got.Name)
	})
}

func TestUser_Update(t *testing.T) {
	test := newUserTest()
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
		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("repo returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "asdasdfasd",
			UserName: "asdasdasdasd",
			Email:    "idzharbae@gmail.com",
			Phone:    "123123123",
			PhotoURL: "asdasd",
			Password: "asdasdasd",
			Type:     1,
		}
		test.writer.EXPECT().Update(entity.User{
			ID:          req.ID,
			Name:        req.Name,
			Phone:       req.Phone,
			PhotoURL:    req.PhotoURL,
			Password:    req.Password,
			NewPassword: req.NewPassword,
			Address:     req.Address,
			Description: req.Description,
		}).Return(entity.User{}, errors.New("error"))

		got, err := test.unit.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("repo returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			ID:       1,
			Name:     "asdasdasdasdas",
			UserName: "asdasdasdasd",
			Email:    "idzharbae@gmail.com",
			Phone:    "123123123",
			PhotoURL: "asdasd",
			Password: "asdasdasd",
			Type:     1,
		}
		test.writer.EXPECT().Update(entity.User{
			ID:          req.ID,
			Name:        req.Name,
			Phone:       req.Phone,
			PhotoURL:    req.PhotoURL,
			Password:    req.Password,
			NewPassword: req.NewPassword,
			Address:     req.Address,
			Description: req.Description,
		}).Return(req, nil)

		got, err := test.unit.Update(req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.User{}, got)
	})
}
