package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/domain"
	"fizz/internal/core/port/outgoing"
	"fizz/internal/pkg/logger"
)

type EmailService struct {
	emailRepository outgoing.EmailRepository
	senderMailgun   outgoing.EmailSender
}

func (e EmailService) Send(ctx context.Context, email dto.SendEmail, backend string) error {
	from := domain.NewEmailAddress(email.From)

	var to []domain.EmailAddress
	for _, v := range email.To {
		to = append(to, domain.NewEmailAddress(v))
	}

	subject := domain.NewSubject(email.Subject)
	body := domain.NewMessageBody(email.Body)

	emailID := e.emailRepository.NextEmailID(ctx)

	emailDomain, err := domain.NewEmail(emailID, from, to, nil, subject, body)
	if err != nil {
		return err
	}

	go func() {
		emailDomain.MarkAsQueued()

		var refID domain.EmailID

		if backend == "MAILGUN" {
			emailDomain.SetEmailBackend(domain.MAILGUN)

			refID, err = e.senderMailgun.Send(ctx, emailDomain)
			if err != nil {
				logger.Log.Error("email queueing error", err)
				return
			}

		} else if backend == "SENDGRID" {
			emailDomain.SetEmailBackend(domain.SENDGRID)

			// TODO: sendgrid backend
		}

		emailDomain.SetReferenceID(refID)

		if err := e.emailRepository.Store(ctx, emailDomain); err != nil {
			logger.Log.Error("email storing error", err)
			return
		}
	}()

	return nil
}

func NewEmailService(sender outgoing.EmailSender, repository outgoing.EmailRepository) *EmailService {
	return &EmailService{senderMailgun: sender, emailRepository: repository}
}
