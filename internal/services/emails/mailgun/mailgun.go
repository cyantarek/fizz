package mailgun

import (
	"context"
	"time"

	mailgungo "github.com/mailgun/mailgun-go"

	"github.com/cyantarek/fizz/internal/domains"
	"github.com/cyantarek/fizz/internal/pkg/logger"
)

type mailgun struct {
	client mailgungo.Mailgun
}

func (m mailgun) GetStats(ctx context.Context) ([]domains.Stats, error) {
	return nil, nil
}

func (m mailgun) Name() string {
	return "MAILGUN"
}

func New(client mailgungo.Mailgun) *mailgun {
	return &mailgun{client: client}
}

func (m mailgun) Send(ctx context.Context, email domains.Email) (domains.ID, error) {
	var to []string

	for _, v := range email.To() {
		to = append(to, v.Address())
	}

	mgEmail := m.client.NewMessage(email.From().Address(), email.Subject().Value(), email.MessageBody().Value(), to...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, id, err := m.client.Send(mgEmail)
	if err != nil {
		logger.Log.Error("email send unsuccessful [resp] [id] -> ", err.Error())
		return domains.NewID(id), err
	}

	logger.Log.Info("email queue successful [resp] [id] -> ", resp, id)

	return domains.NewID(id), nil
}
