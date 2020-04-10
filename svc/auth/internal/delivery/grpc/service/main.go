package service

import "github.com/idzharbae/marketplace-backend/svc/auth/internal/app"

type Services struct {
	*AuthService
	*UserService
}

func GetServices(a *app.Auth) *Services {
	return &Services{
		AuthService: NewAuthService(a.UseCases.TokenUC, a.UseCases.UserUC),
		UserService: NewUserService(a.UseCases.UserUC),
	}
}
