package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
)

type File struct {
	FileWriter internal.FileWriter
}

func NewFile(writer internal.FileWriter) *File {
	return &File{FileWriter: writer}
}

func (f *File) UploadFile(req entity.File) (entity.File, error) {
	return entity.File{}, nil
}

func (f *File) DeleteFile(req entity.File) error {
	return nil
}
