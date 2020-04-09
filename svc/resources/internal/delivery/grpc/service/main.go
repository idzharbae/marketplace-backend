package service

type Services struct {
	*FileService
}

func GetServices(app *app.Resources) *Services {
	return &Services{NewFileService()}
}
