package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/domain"
	"fizz/internal/core/port/outgoing"
)

type EmailService struct {
	sender outgoing.EmailSender
}

func (e EmailService) Send(ctx context.Context, email dto.SendEmail) error {
	from := domain.NewEmailAddress(email.From)

	var to []domain.EmailAddress
	for _, v := range email.To {
		to = append(to, domain.NewEmailAddress(v))
	}

	subject := domain.NewSubject(email.Subject)
	body := domain.NewMessageBody(email.Body)

	emailDomain := domain.NewEmail(from, to, nil, subject, body)

	return e.sender.Send(ctx, emailDomain)
}

func NewEmailService(sender outgoing.EmailSender) *EmailService {
	return &EmailService{sender: sender}
}
