package fileio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIOUtil_CreateFile(t *testing.T) {
	io := NewIO()
	err := io.CreateFile("/img/test.gif", []byte("GIF89a"))
	assert.Nil(t, err)
}

func TestIOUtil_DeleteFile(t *testing.T) {
	t.Run("file exists, success", func(t *testing.T) {
		io := NewIO()
		err := io.DeleteFile("/img/test.gif")
		assert.Nil(t, err)
	})
	t.Run("file doesn't exist, error", func(t *testing.T) {
		io := NewIO()
		err := io.DeleteFile("/img/asdasd.gif")
		assert.NotNil(t, err)
	})
}
