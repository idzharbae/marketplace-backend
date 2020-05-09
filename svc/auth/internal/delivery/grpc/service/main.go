package service

import "github.com/idzharbae/marketplace-backend/svc/auth/internal/app"

type Services struct {
	*AuthService
	*UserService
	*PaymentService
	*SaldoHistoryService
}

func GetServices(a *app.Auth) *Services {
	return &Services{
		AuthService:         NewAuthService(a.UseCases.TokenUC, a.UseCases.UserUC),
		UserService:         NewUserService(a.UseCases.UserUC),
		PaymentService:      NewPaymentService(a.UseCases.PaymentUC),
		SaldoHistoryService: NewSaldoHistoryService(a.UseCases.SaldoHistoryUC),
	}
}
