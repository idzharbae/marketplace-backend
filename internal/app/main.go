package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Grpc struct{
		Address string
	}
}

type Marketplace struct {
	Config Config
	UseCases *UseCases
	Repo *Repo
}

func NewMarketplace(cfgPath string) *Marketplace {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		panic(err)
	}
	UC := NewUsecase()
	Repo := NewRepo()

	return &Marketplace{
		Config:   cfg,
		UseCases: UC,
		Repo:     Repo,
	}
}

func readConfig(filepath string) (Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	var cfg Config
	err = json.Unmarshal(bytes, cfg)
	return cfg, err
}