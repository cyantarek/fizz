package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/domain"
	"fizz/internal/core/port/outgoing"
	"fizz/internal/pkg/logger"
	"time"
)

type EmailSender map[string]outgoing.EmailSender

type EmailService struct {
	emailRepository outgoing.EmailRepository
	senders         EmailSender
}

func (e EmailService) LookupStatus(ctx context.Context, id string) (*dto.LookupEmail, error) {
	emailDomain, err := e.emailRepository.LookupStatus(ctx, domain.NewID(id))
	if err != nil {
		return nil, err
	}

	var out dto.LookupEmail

	out.ID = emailDomain.Id().String()
	out.From = emailDomain.From().Address()
	out.EmailBackend = string(emailDomain.EmailBackend())
	out.Status = emailDomain.Status().String()

	for _, addr := range emailDomain.To() {
		out.To = append(out.To, addr.Address())
	}

	return &out, nil
}

func (e EmailService) Send(ctx context.Context, email dto.SendEmail, backend string) error {
	// build the domain model from dto
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

	emailDomain.SetEmailBackend(domain.EmailBackend(backend))

	// queue email asynchronously
	go func() {
		emailDomain.MarkAsQueued()

		var refID domain.ID

		// retry with exponential backoff
		for i := 1; i <= 5; i++ {
			refID, err = e.senders[backend].Send(ctx, emailDomain)
			if err != nil {
				logger.Log.Error("email queueing error", err)

				time.Sleep(time.Second * time.Duration(i))
				continue
			}

			logger.Log.Info("Email sent successfully")
			break
		}

		emailDomain.SetReferenceID(refID)

		// store to db
		if err := e.emailRepository.Store(context.Background(), emailDomain); err != nil {
			logger.Log.Error("email storing error ", err)
			return
		}
	}()

	logger.Log.Info("Email queued successfully")

	return nil
}

func NewEmailService(sender outgoing.EmailSender, repository outgoing.EmailRepository) *EmailService {
	return &EmailService{senders: map[string]outgoing.EmailSender{
		sender.Name(): sender,
	}, emailRepository: repository}
}
