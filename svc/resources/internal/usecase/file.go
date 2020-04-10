package usecase

import (
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/resources/constants"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/resources/util"
)

type File struct {
	FileWriter      internal.FileWriter
	OwnershipWriter internal.OwnershipWriter
	OwnershipReader internal.OwnershipReader
}

func NewFile(fileWriter internal.FileWriter, ownershipWriter internal.OwnershipWriter, reader internal.OwnershipReader) *File {
	return &File{FileWriter: fileWriter, OwnershipWriter: ownershipWriter, OwnershipReader: reader}
}

func (f *File) UploadFile(req entity.File) (entity.File, error) {
	err := f.validateInput(req)
	if err != nil {
		return entity.File{}, err
	}

	res, err := f.FileWriter.UploadFile(req)
	if err != nil {
		return entity.File{}, err
	}

	_, err = f.OwnershipWriter.Save(res)
	if err != nil {
		return entity.File{}, err
	}

	return res, nil
}

func (f *File) DeleteFile(req entity.File) error {
	file, err := f.OwnershipReader.GetByURL(req.URL)
	if err != nil {
		return err
	}
	if req.OwnerID != file.OwnerID {
		return errors.New("user is not authorized to delete this file")
	}

	err = f.OwnershipWriter.DeleteByID(file.ID)
	if err != nil {
		return err
	}
	return f.FileWriter.DeleteFile(file)
}

func (f *File) validateInput(req entity.File) error {
	if req.OwnerID == 0 {
		return errors.New("owner id should not be empty")
	}
	if len(req.Data) > constants.MaxFileSize {
		return errors.New("file size is too large")
	}
	if !util.IsValidFileExt(req.Extension) {
		return errors.New("invalid file extension")
	}
	if !util.IsImage(req.Data) {
		return errors.New("invalid file type")
	}
	return nil
}
