package repo

import (
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/svc/catalog/util/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ProductWriterTest struct {
	DB   *gormmock.MockGormw
	Ctrl *gomock.Controller
	Unit *ProductWriter
	Req  entity.Product
}

func NewProductWriterTest() *ProductWriterTest {
	return &ProductWriterTest{}
}

func (p *ProductWriterTest) Begin(t *testing.T) {
	p.Ctrl = gomock.NewController(t)
	p.DB = gormmock.NewMockGormw(p.Ctrl)
	p.Unit = NewProductWriter(p.DB)
	p.Req = entity.Product{ID: 23, Name: "name", Slug: "test"}
}

func (p *ProductWriterTest) Finish() {
	p.Ctrl.Finish()
}

func TestProductWriter_Create(t *testing.T) {
	test := NewProductWriterTest()
	t.Run("connection returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Create(test.Req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("connection returns no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.Create(test.Req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.Product{}, got)
	})
	t.Run("given input with ID, should ignore the ID", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Save(gomock.Any()).DoAndReturn(func(value *model.Product) connection.Gormw {
			assert.Equal(t, int32(0), value.ID)
			*value = model.Product{
				ID:   1,
				Name: "test",
			}
			return test.DB
		})
		test.DB.EXPECT().Error().Return(nil)
		got, err := test.Unit.Create(test.Req)
		assert.Nil(t, err)
		assert.Equal(t, int32(1), got.ID)
	})
}

func TestProductWriter_Update(t *testing.T) {
	test := NewProductWriterTest()
	t.Run("connection returns error when finding, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Update(test.Req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("product does not exists, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(true)

		got, err := test.Unit.Update(test.Req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("error when saving, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		got, err := test.Unit.Update(test.Req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Product{}, got)
	})
	t.Run("success when saving, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Save(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)

		got, err := test.Unit.Update(test.Req)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.Product{}, got)
	})
}

func TestProductWriter_DeleteByID(t *testing.T) {
	test := NewProductWriterTest()
	t.Run("error when finding, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		err := test.Unit.DeleteByID(test.Req.ID)
		assert.NotNil(t, err)
	})
	t.Run("record not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(true)

		err := test.Unit.DeleteByID(test.Req.ID)
		assert.NotNil(t, err)
	})
	t.Run("error when deleting, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Delete(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		err := test.Unit.DeleteByID(test.Req.ID)
		assert.NotNil(t, err)
	})
	t.Run("no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Delete(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)

		err := test.Unit.DeleteByID(test.Req.ID)
		assert.Nil(t, err)
	})
}

func TestProductWriter_DeleteBySlug(t *testing.T) {
	test := NewProductWriterTest()
	t.Run("error when finding, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		err := test.Unit.DeleteBySlug(test.Req.Slug)
		assert.NotNil(t, err)
	})
	t.Run("record not found, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(true)

		err := test.Unit.DeleteBySlug(test.Req.Slug)
		assert.NotNil(t, err)
	})
	t.Run("error when deleting, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Delete(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(errors.New("error"))

		err := test.Unit.DeleteBySlug(test.Req.Slug)
		assert.NotNil(t, err)
	})
	t.Run("no error, should return no error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()

		test.DB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(test.DB)
		test.DB.EXPECT().First(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)
		test.DB.EXPECT().RecordNotFound().Return(false)
		test.DB.EXPECT().Delete(gomock.Any()).Return(test.DB)
		test.DB.EXPECT().Error().Return(nil)

		err := test.Unit.DeleteBySlug(test.Req.Slug)
		assert.Nil(t, err)
	})
}
