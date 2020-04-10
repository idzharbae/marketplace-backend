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

type ownershipReaderTest struct {
	ctrl   *gomock.Controller
	mockdb *gormmock.MockGormw
	unit   internal.OwnershipReader
}

func newOwnershipReaderTest() *ownershipReaderTest {
	return &ownershipReaderTest{}
}

func (ort *ownershipReaderTest) Begin(t *testing.T) {
	ort.ctrl = gomock.NewController(t)
	ort.mockdb = gormmock.NewMockGormw(ort.ctrl)
	ort.unit = NewOwnershipReader(ort.mockdb)
}

func (ort *ownershipReaderTest) Finish() {
	ort.ctrl.Finish()
}

func TestOwnershipReader_GetByURL(t *testing.T) {
	test := newOwnershipReaderTest()
	t.Run("record not found should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdf"
		test.mockdb.EXPECT().Where("file_url=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).Return(test.mockdb)
		test.mockdb.EXPECT().RecordNotFound().Return(true)

		got, err := test.unit.GetByURL(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("db returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdf"
		resp := model.FileOwnership{
			ID:       123,
			FilePath: "asdf",
			FileURL:  "asdf",
			OwnerID:  21,
		}
		test.mockdb.EXPECT().Where("file_url=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			*arg = resp
			return test.mockdb
		})
		test.mockdb.EXPECT().RecordNotFound().Return(false)
		test.mockdb.EXPECT().Error().Return(errors.New("error"))

		got, err := test.unit.GetByURL(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("success", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := "asdf"
		resp := model.FileOwnership{
			ID:       123,
			FilePath: "/pdf/asdf.jpg",
			FileURL:  "http://127.0.0.1:1441/pdf/asdf.jpg",
			OwnerID:  21,
		}
		test.mockdb.EXPECT().Where("file_url=?", req).Return(test.mockdb)
		test.mockdb.EXPECT().First(gomock.Any()).DoAndReturn(func(arg *model.FileOwnership) *gormmock.MockGormw {
			*arg = resp
			return test.mockdb
		})
		test.mockdb.EXPECT().RecordNotFound().Return(false)
		test.mockdb.EXPECT().Error().Return(nil)

		got, err := test.unit.GetByURL(req)
		assert.Nil(t, err)
		assert.Equal(t, entity.File{
			ID:        resp.ID,
			OwnerID:   resp.OwnerID,
			Name:      "asdf",
			Extension: "jpg",
			Type:      "pdf",
			URL:       resp.FileURL,
		}, got)
	})
}
