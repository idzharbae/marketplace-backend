package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/idzharbae/marketplace-backend/internal/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	flag.Usage = usage
	flag.Parse()

	dbCfg, err := readConfig(basepath + "/../config/marketplace.json")
	if err != nil {
		log.Fatal(err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", dbCfg.Db.Master.Address, dbCfg.Db.Master.Port),
		Database: dbCfg.Db.Master.DbName,
		User:     dbCfg.Db.Master.User,
		Password: dbCfg.Db.Master.Password,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func readConfig(filepath string) (config.Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return config.Config{}, err
	}
	if jsonFile == nil {
		return config.Config{}, errors.New("couldnt read path: " + filepath)
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return config.Config{}, err
	}
	if bytes == nil {
		log.Fatal(bytes)
	}
	var cfg config.Config
	err = json.Unmarshal([]byte(bytes), &cfg)
	if err != nil {
		log.Printf("%v", bytes)
		return config.Config{}, err
	}
	return cfg, nil
}
