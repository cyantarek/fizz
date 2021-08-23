package driven

import (
	"context"
	"fizz/internal/core/domain"
	"fizz/internal/pkg/logger"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

type MailgunEmail struct {
	client mailgun.Mailgun
}

func (m MailgunEmail) GetStats(ctx context.Context) ([]domain.Stats, error) {
	evs := []string{
		"accepted",
		"delivered",
		"failed",
		"opened",
		"clicked",
		"unsubscribed",
		"complained",
		"stored",
	}

	stats, err := m.client.GetStats(ctx, evs, &mailgun.GetStatOptions{
		Resolution: "month",
		Duration:   "1m",
	})
	if err != nil {
		return nil, err
	}

	var out []domain.Stats

	for _, stat := range stats {
		out = append(out, domain.Stats{
			Accepted:     stat.Accepted.Total,
			Delivered:    stat.Delivered.Total,
			Failed:       stat.Failed.Permanent.Total,
			Stored:       stat.Stored.Total,
			Opened:       stat.Opened.Total,
			Clicked:      stat.Clicked.Total,
			Unsubscribed: stat.Unsubscribed.Total,
			Complained:   stat.Complained.Total,
		})
	}

	return out, err
}

func (m MailgunEmail) Name() string {
	return "MAILGUN"
}

func NewMailgunEmail(client mailgun.Mailgun) *MailgunEmail {
	return &MailgunEmail{client: client}
}

func (m MailgunEmail) Send(ctx context.Context, email domain.Email) (domain.ID, error) {
	var to []string

	for _, v := range email.To() {
		to = append(to, v.Address())
	}

	mgEmail := m.client.NewMessage(email.From().Address(), email.Subject().Value(), email.MessageBody().Value(), to...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, id, err := m.client.Send(ctx, mgEmail)
	if err != nil {
		logger.Log.Error("email send unsuccessful [resp] [id] -> ", err.Error())
		return domain.NewID(id), err
	}

	logger.Log.Info("email queue successful [resp] [id] -> ", resp, id)

	return domain.NewID(id), nil
}
