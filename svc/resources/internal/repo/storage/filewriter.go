package storage

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
)

type FileWriter struct {
	IO internal.FileIO
}

func NewFileWriter(io internal.FileIO) *FileWriter {
	return &FileWriter{
		IO: io,
	}
}

func (fw *FileWriter) UploadFile(req entity.File) (entity.File, error) {
	return entity.File{}, nil
}

func (fw *FileWriter) DeleteFile(req entity.File) error {
	return nil
}
