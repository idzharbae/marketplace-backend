package fileio

import (
	"os"
	"path/filepath"
	"runtime"
)

type IOUtil struct {
}

func NewIO() *IOUtil {
	return &IOUtil{}
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func (io *IOUtil) CreateFile(path string, data []byte) error {
	file, err := os.Create(basepath + "/../../../public" + path)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (io *IOUtil) DeleteFile(path string) error {
	return os.Remove(basepath + "/../../../public" + path)
}
