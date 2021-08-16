package driven

import (
	"context"
	"fizz/internal/core/domain"
	"fizz/internal/pkg/logger"
	"github.com/mailgun/mailgun-go/v3"
)

type MailgunEmail struct {
	client mailgun.Mailgun
}

func NewMailgunEmail(client mailgun.Mailgun) *MailgunEmail {
	return &MailgunEmail{client: client}
}

func (m MailgunEmail) Send(ctx context.Context, email domain.Email) error {
	logger.Log.Println("email sent")

	return nil
}
