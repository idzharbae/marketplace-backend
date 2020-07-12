package gateway

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmtpMailer_SendMail(t *testing.T) {
	t.Skip() // Integration test, disabled by default
	t.Run("should successfully send email", func(t *testing.T) {
		var password string
		var email string
		var destination string

		mailer := NewSmtpMailer(email, password, "smtp.gmail.com", "587")
		err := mailer.SendMail(requests.SendMailReq{
			Destination: destination,
			Title:       "Test email title",
			Message:     "test email message",
		})
		assert.Nil(t, err)
	})
}
