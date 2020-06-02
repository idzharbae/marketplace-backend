package storage

import (
	"github.com/idzharbae/marketplace-backend/svc/resources/internal"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	"strconv"
	"time"
)

type FileWriter struct {
	cfg config.Config
	IO  internal.FileIO
}

func NewFileWriter(io internal.FileIO, cfg config.Config) *FileWriter {
	return &FileWriter{
		IO:  io,
		cfg: cfg,
	}
}

func (fw *FileWriter) UploadFile(req entity.File) (entity.File, error) {
	fileName := strconv.Itoa(int(time.Now().Unix())) + "." + req.Extension
	err := fw.IO.CreateFile("/img/"+fileName, req.Data)
	if err != nil {
		return entity.File{}, err
	}
	return entity.File{
		OwnerID:   req.OwnerID,
		Name:      fileName,
		Type:      "img",
		Extension: req.Extension,
		URL:       "http://" + fw.cfg.REST.IP + fw.cfg.REST.Port + "/img/" + fileName,
	}, nil
}

func (fw *FileWriter) DeleteFile(req entity.File) error {
	fileName := req.Name + "." + req.Extension
	return fw.IO.DeleteFile("/" + req.Type + "/" + fileName)
}
