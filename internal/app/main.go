package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Grpc struct {
		Port string
	}
}

type Marketplace struct {
	Config   Config
	UseCases *UseCases
	Repos    *Repos
}

func NewMarketplace(cfgPath string) (*Marketplace, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	UCS := NewUsecase()
	Repos := NewRepos()

	return &Marketplace{
		Config:   cfg,
		UseCases: UCS,
		Repos:    Repos,
	}, nil
}

func readConfig(filepath string) (Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	if jsonFile == nil {
		return Config{}, errors.New("couldnt read path: " + filepath)
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}
	if bytes == nil {
		log.Fatal(bytes)
	}
	var cfg Config
	err = json.Unmarshal([]byte(bytes), &cfg)
	if err != nil {
		log.Printf("%v", bytes)
		return Config{}, err
	}
	return cfg, nil
}

func (m *Marketplace) Close() []error {
	var errs []error

	errs = append(errs, m.Repos.Close()...)
	errs = append(errs, m.UseCases.Close()...)

	return errs
}
