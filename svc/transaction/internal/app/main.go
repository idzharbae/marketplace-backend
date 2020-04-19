package app

import (
	"encoding/json"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/transaction/internal/config"
	"io/ioutil"
	"log"
	"os"
)

type Transaction struct {
	Config   config.Config
	UseCases *UseCases
	Repos    *Repos
	Gateways *Gateways
}

func NewTransaction(cfgPath string) (*Transaction, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	repos := NewRepos(cfg)
	gateways, err := NewGateways(cfg)
	if err != nil {
		return nil, err
	}
	usecases := NewUseCases(repos, gateways)

	return &Transaction{
		Config:   cfg,
		Repos:    repos,
		UseCases: usecases,
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

func (t *Transaction) Close() []error {
	var errs []error

	errs = append(errs, t.Repos.Close()...)
	errs = append(errs, t.UseCases.Close()...)

	return errs
}
