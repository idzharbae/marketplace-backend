package repo

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userReaderTest struct {
	ctrl *gomock.Controller
	db   *gormmock.MockGormw
	unit internal.UserReader
}

func newUserReaderTest() *userReaderTest {
	return &userReaderTest{}
}

func (rt *userReaderTest) Begin(t *testing.T) {
	rt.ctrl = gomock.NewController(t)
	rt.db = gormmock.NewMockGormw(rt.ctrl)
	rt.unit = NewUserReader(rt.db)
}

func (rt *userReaderTest) Finish() {
	rt.ctrl.Finish()
}

func TestUserReader_GetByEmailAndPassword(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			Email:    "asdf@asdf.com",
			Password: "laksdlaksd",
		}

		test.db.EXPECT().Where("email=?", req.Email).Return(test.db)
		test.db.EXPECT().Where("password=?", fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByEmailAndPassword(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			Email:    "asdf@asdf.com",
			Password: "laksdlaksd",
		}

		test.db.EXPECT().Where("email=?", req.Email).Return(test.db)
		test.db.EXPECT().Where("password=?", fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.UserFromEntity(req)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByEmailAndPassword(req)
		assert.Nil(t, err)
		assert.Equal(t, req.ID, got.ID)
		assert.Equal(t, req.Name, got.Name)
		assert.Equal(t, req.UserName, got.UserName)
		assert.Equal(t, req.Email, got.Email)
	})
}

func TestUserReader_GetByUserNameAndPassword(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			UserName: "asdasdasd",
			Password: "laksdlaksd",
		}

		test.db.EXPECT().Where("user_name=?", req.UserName).Return(test.db)
		test.db.EXPECT().Where("password=?", fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByUserNameAndPassword(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.User{
			UserName: "asdasd",
			Password: "laksdlaksd",
		}

		test.db.EXPECT().Where("user_name=?", req.UserName).Return(test.db)
		test.db.EXPECT().Where("password=?", fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gormmock.MockGormw {
			*user = model.UserFromEntity(req)
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByUserNameAndPassword(req)
		assert.Nil(t, err)
		assert.Equal(t, req.ID, got.ID)
		assert.Equal(t, req.Name, got.Name)
		assert.Equal(t, req.UserName, got.UserName)
		assert.Equal(t, req.Email, got.Email)
	})
}

func TestUserReader_GetByID(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1234)
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByID(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1234)
		resp := model.User{
			ID:       1234,
			Name:     "asdf",
			UserName: "asdasd",
			Email:    "asdasd",
			Phone:    "asdasd",
			Password: "asdasd",
			PhotoURL: "Asdasd",
		}
		test.db.EXPECT().Where("id=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = resp
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByID(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.UserName, got.UserName)
		assert.Equal(t, resp.Email, got.Email)
		assert.Equal(t, resp.PhotoURL, got.PhotoURL)
		assert.Equal(t, "", got.Password)
	})
}

func TestUserReader_GetByUserName(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdfg"
		test.db.EXPECT().Where("user_name=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByUserName(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdfg"
		resp := model.User{
			ID:       1234,
			Name:     "asdf",
			UserName: "asdasd",
			Email:    "asdasd",
			Phone:    "asdasd",
			Password: "asdasd",
			PhotoURL: "Asdasd",
		}

		test.db.EXPECT().Where("user_name=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = resp
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByUserName(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.UserName, got.UserName)
		assert.Equal(t, resp.Email, got.Email)
		assert.Equal(t, resp.PhotoURL, got.PhotoURL)
		assert.Equal(t, "", got.Password)
	})
}

func TestUserReader_GetByEmail(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db return error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdfg"
		test.db.EXPECT().Where("email=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByEmail(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.User{}, got)
	})
	t.Run("db return no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdfg"
		resp := model.User{
			ID:       1234,
			Name:     "asdf",
			UserName: "asdasd",
			Email:    "asdasd",
			Phone:    "asdasd",
			Password: "asdasd",
			PhotoURL: "Asdasd",
		}

		test.db.EXPECT().Where("email=?", req).Return(test.db)
		test.db.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.User) *gormmock.MockGormw {
			*arg = resp
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByEmail(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.UserName, got.UserName)
		assert.Equal(t, resp.Email, got.Email)
		assert.Equal(t, resp.PhotoURL, got.PhotoURL)
		assert.Equal(t, "", got.Password)
	})
}

func TestUserReader_GetShopsByProvince(t *testing.T) {
	test := newUserReaderTest()
	t.Run("db returns error should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		test.db.EXPECT().Where("province=?", "asdf").Return(test.db)
		test.db.EXPECT().Where("type=?", globalconstant.ShopType).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).Return(test.db)
		test.db.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetShopsByProvince("asdf")
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})
	t.Run("db returns no error should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		test.db.EXPECT().Where("province=?", "asdf").Return(test.db)
		test.db.EXPECT().Where("type=?", globalconstant.ShopType).Return(test.db)
		test.db.EXPECT().Find(gomock.Any()).DoAndReturn(func(arg *[]model.User) *gormmock.MockGormw {
			*arg = append(*arg, model.User{ID: 1})
			return test.db
		})
		test.db.EXPECT().Error().Return(nil)

		got, err := test.unit.GetShopsByProvince("asdf")
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
