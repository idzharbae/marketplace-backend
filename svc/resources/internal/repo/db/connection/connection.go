package connection

import (
	"fmt"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type ConnectionParams struct {
	dbEngine string
	host     string
	port     int
	username string
	password string
	dbName   string
	sslMode  string
}

func NewConnection(dBEngine string, args interface{}) (Gormw, error) {
	if dBEngine == "" {
		dBEngine = "postgres"
	}

	db, err := Openw(dBEngine, args)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create db connection")
	}
	return db, nil
}

func GetConnectionParams(param config.DbParams) string {
	if param.SSLMode == "" {
		param.SSLMode = "disable"
	}
	args := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		param.Address,
		param.Port,
		param.User,
		param.DbName,
		param.Password,
		param.SSLMode)
	return args
}
