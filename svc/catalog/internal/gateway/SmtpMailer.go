package gateway

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"net/smtp"
)

type SmtpMailer struct {
	source      string
	password    string
	smtpAddress string
	smtpPort    string
}

func NewSmtpMailer(source, password, smtpAddress, smtpPort string) *SmtpMailer {
	return &SmtpMailer{
		source:      source,
		password:    password,
		smtpAddress: smtpAddress,
		smtpPort:    smtpPort,
	}
}

func (sm *SmtpMailer) SendMail(req requests.SendMailReq) error {
	msg := "From: " + sm.source + "\n" +
		"To: " + req.Destination + "\n" +
		"Subject: " + req.Title + "\n\n" +
		req.Message

	return smtp.SendMail(sm.smtpAddress+":"+sm.smtpPort,
		smtp.PlainAuth("", sm.source, sm.password, sm.smtpAddress),
		sm.source, []string{req.Destination}, []byte(msg))
}
