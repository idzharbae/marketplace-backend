package repo

import (
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/internal/entity"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/internal/repo/connection/gormmock"
	"github.com/idzharbae/marketplace-backend/internal/repo/model"
	"github.com/idzharbae/marketplace-backend/util/errors"
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
	p.Req = entity.Product{ID: 23, Name: "test"}
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
