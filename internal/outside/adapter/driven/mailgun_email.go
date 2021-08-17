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
	var to []string

	for _, v := range email.To() {
		to = append(to, v.Address())
	}

	mgEmail := m.client.NewMessage(email.From().Address(), email.Subject().Value(), email.MessageBody().Value(), to...)

	go func() {
		resp, id, err := m.client.Send(ctx, mgEmail)
		if err != nil {
			logger.Log.Error("email send unsuccessful [resp] [id] -> ", err.Error())
		}

		logger.Log.Info("email send successfully [resp] [id] -> ", resp, id)
	}()

	return nil
}
