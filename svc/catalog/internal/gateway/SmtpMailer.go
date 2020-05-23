package gateway

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"net/smtp"
)

type SmtpMailer struct {
	source      string
	password    string
	smtpAddress string
}

func NewSmtpMailer(source, password, smtpAddress string) *SmtpMailer {
	return &SmtpMailer{source: source, password: password, smtpAddress: smtpAddress}
}

func (sm *SmtpMailer) SendMail(req requests.SendMailReq) error {
	msg := "From: " + req.Destination + "\n" +
		"To: " + sm.source + "\n" +
		"Subject: " + req.Title + "\n\n" +
		req.Message

	return smtp.SendMail(sm.smtpAddress,
		smtp.PlainAuth("", sm.source, sm.password, sm.smtpAddress),
		sm.source, []string{req.Destination}, []byte(msg))
}
