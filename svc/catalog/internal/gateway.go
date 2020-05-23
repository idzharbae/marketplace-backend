package internal

import "github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"

type Mailer interface {
	SendMail(req requests.SendMailReq) error
}
