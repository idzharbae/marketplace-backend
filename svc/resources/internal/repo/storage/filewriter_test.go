package storage

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/bridge/bridgemock"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fileWriterTest struct {
	ctrl *gomock.Controller
	io   *bridgemock.MockFileIO
	unit internal.FileWriter
}

func newFileWriterTest() *fileWriterTest {
	return &fileWriterTest{}
}

func (fwt *fileWriterTest) Begin(t *testing.T) {
	fwt.ctrl = gomock.NewController(t)
	fwt.io = bridgemock.NewMockFileIO(fwt.ctrl)
	fwt.unit = NewFileWriter(fwt.io, config.Config{
		REST: struct {
			IP   string
			Port string
		}{
			IP:   "123.123.123.123",
			Port: ":1441",
		},
	})
}

func (fwt *fileWriterTest) Finish() {
	fwt.ctrl.Finish()
}

func TestFileWriter_UploadFile(t *testing.T) {
	test := newFileWriterTest()
	t.Run("io returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			Extension: "jpg",
			Data:      []byte("GIF89a"),
		}
		test.io.EXPECT().CreateFile(gomock.Any(), gomock.Any()).DoAndReturn(func(fileName string, data []byte) error {
			assert.Equal(t, "/img/", fileName[:5])
			assert.Equal(t, len(req.Data), len(data))

			return errors.New("error")
		})

		got, err := test.unit.UploadFile(req)
		assert.NotNil(t, err)
		assert.Equal(t, entity.File{}, got)
	})
	t.Run("io returns no error, should not return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			OwnerID:   1234,
			Extension: "jpg",
			Data:      []byte("GIF89a"),
		}
		test.io.EXPECT().CreateFile(gomock.Any(), gomock.Any()).DoAndReturn(func(fileName string, data []byte) error {
			assert.Equal(t, "/img/", fileName[:5])
			assert.Equal(t, len(req.Data), len(data))

			return nil
		})

		got, err := test.unit.UploadFile(req)
		assert.Nil(t, err)
		assert.Equal(t, "http://123.123.123.123:1441/img/", got.URL[:32])
		assert.Equal(t, ".jpg", got.URL[len(got.URL)-4:])
		assert.Equal(t, "img", got.Type)
		assert.Equal(t, req.OwnerID, got.OwnerID)
	})
}

func TestFileWriter_DeleteFile(t *testing.T) {
	test := newFileWriterTest()
	t.Run("io returns error, should return error", func(t *testing.T) {
		test.Begin(t)
		defer test.Finish()
		req := entity.File{
			ID:        123,
			Name:      "asdf",
			Extension: "jpg",
			Type:      "img",
			URL:       "asdas",
			OwnerID:   1337,
		}
		test.io.EXPECT().DeleteFile("/img/asdf.jpg").Return(errors.New("error"))

		err := test.unit.DeleteFile(req)
		assert.NotNil(t, err)
	})
}
