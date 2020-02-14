package app

import (
	"encoding/json"
	"errors"
	"github.com/idzharbae/marketplace-backend/internal/config"
	"io/ioutil"
	"log"
	"os"
)

type Marketplace struct {
	Config   config.Config
	UseCases *UseCases
	Repos    *Repos
}

func NewMarketplace(cfgPath string) (*Marketplace, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	Repos := NewRepos(cfg)
	UCS := NewUsecase(Repos)

	return &Marketplace{
		Config:   cfg,
		UseCases: UCS,
		Repos:    Repos,
	}, nil
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

func (m *Marketplace) Close() []error {
	var errs []error

	errs = append(errs, m.Repos.Close()...)
	errs = append(errs, m.UseCases.Close()...)

	return errs
}
