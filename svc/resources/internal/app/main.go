package app

import (
	"encoding/json"
	"errors"
	"github.com/idzharbae/marketplace-backend/svc/resources/internal/config"
	"io/ioutil"
	"log"
	"os"
)

type Resources struct {
	Config   config.Config
	UseCases *Usecases
	Repos    *Repos
	Bridges  *Bridges
}

func NewResources(cfgPath string) (*Resources, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	Bridges := NewBridges()
	Repos := NewRepos(Bridges)
	UCS := NewUsecases(Repos)

	return &Resources{
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

func (r *Resources) Close() []error {
	var errs []error

	errs = append(errs, r.Repos.Close()...)
	errs = append(errs, r.UseCases.Close()...)

	return errs
}
