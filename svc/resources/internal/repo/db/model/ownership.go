package model

import (
	"strings"
	"time"
)

type FileOwnership struct {
	ID        int64
	FilePath  string
	FileURL   string
	OwnerID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f FileOwnership) TableName() string {
	return "file_ownership"
}

func (f FileOwnership) Name() string {
	nameAndExt := strings.Split(f.FilePath, "/")[2]
	name := strings.Split(nameAndExt, ".")[0]
	return name
}

func (f FileOwnership) Ext() string {
	nameAndExt := strings.Split(f.FilePath, "/")[2]
	return strings.Split(nameAndExt, ".")[1]
}

func (f FileOwnership) Type() string {
	return strings.Split(f.FilePath, "/")[1]
}
