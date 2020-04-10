package db

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/repo/db/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ownershipWriterTest struct {
	ctrl   *gomock.Controller
	mockdb *gormmock.MockGormw
	unit   internal.OwnershipWriter
}

func newOwnershipWriterTest() *ownershipWriterTest {
	return &ownershipWriterTest{}
}

func (owt *ownershipWriterTest) Begin(t *testing.T) {
	owt.ctrl = gomock.NewController(t)
	owt.mockdb = gormmock.NewMockGormw(owt.ctrl)
	owt.unit = NewOwnershipWriter(owt.mockdb)
}

func (owt *ownershipWriterTest) Finish() {
	owt.ctrl.Finish()
}

func TestOwnershipWriter_Save(t *testing.T) {
	test := newOwnershipWriterTest()
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			Name:      "test",
			Extension: "Test",
			Type:      "test",
			URL:       "Test",
			OwnerID:   123,
		}
		test.mockdb.EXPECT().Save(gomock.Any()).Return(test.mockdb)
		test.mockdb.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.Save(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			Name:      "test",
			Extension: "Test",
			Type:      "test",
			URL:       "Test",
			OwnerID:   123,
		}
		test.mockdb.EXPECT().Save(gomock.Any()).Return(test.mockdb)
		test.mockdb.EXPECT().Error().Return(nil)

		got, err := test.unit.Save(req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.File{}, got)
	})
}

func TestOwnershipWriter_Delete(t *testing.T) {
	test := newOwnershipWriterTest()
	t.Run("file not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1337)
		test.mockdb.EXPECT().Where("id=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).Return(test.mockdb)
		test.mockdb.EXPECT().RecordNotFound().Return(true)

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
	t.Run("error when querying, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1337)
		resp := model.FileOwnership{
			ID:       1337,
			FilePath: "asdf",
			FileURL:  "Asdf",
			OwnerID:  13,
		}
		test.mockdb.EXPECT().Where("id=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			*arg = resp
			return test.mockdb
		})
		test.mockdb.EXPECT().RecordNotFound().Return(false)
		test.mockdb.EXPECT().Error().Return(errors.New("error"))

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
	t.Run("error when saving, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1337)
		resp := model.FileOwnership{
			ID:       1337,
			FilePath: "asdf",
			FileURL:  "Asdf",
			OwnerID:  13,
		}
		test.mockdb.EXPECT().Where("id=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			*arg = resp
			return test.mockdb
		})
		test.mockdb.EXPECT().RecordNotFound().Return(false)
		test.mockdb.EXPECT().Error().Return(nil)
		test.mockdb.EXPECT().Delete(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			assert.Equal(t, *arg, resp)
			return test.mockdb
		})
		test.mockdb.EXPECT().Error().Return(errors.New("error"))

		err := test.unit.DeleteByID(req)
		assert.NotNil(t, err)
	})
	t.Run("success when saving, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := int64(1337)
		resp := model.FileOwnership{
			ID:       1337,
			FilePath: "asdf",
			FileURL:  "Asdf",
			OwnerID:  13,
		}
		test.mockdb.EXPECT().Where("id=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			*arg = resp
			return test.mockdb
		})
		test.mockdb.EXPECT().RecordNotFound().Return(false)
		test.mockdb.EXPECT().Error().Return(nil)
		test.mockdb.EXPECT().Delete(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			assert.Equal(t, *arg, resp)
			return test.mockdb
		})
		test.mockdb.EXPECT().Error().Return(nil)

		err := test.unit.DeleteByID(req)
		assert.Nil(t, err)
	})
}
