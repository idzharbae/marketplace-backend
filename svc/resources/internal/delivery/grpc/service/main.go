package service

import "github.com/idzharbae/marketplace-backend/svc/resources/internal/app"

type Services struct {
	*FileService
}

func GetServices(App *app.Resources) *Services {
	return &Services{NewFileService(App.UseCases.File)}
}
