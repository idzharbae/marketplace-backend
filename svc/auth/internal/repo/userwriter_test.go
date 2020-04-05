package repo

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/repo/model"
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
