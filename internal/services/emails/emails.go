package emails

import (
	"context"
	"github.com/cyantarek/fizz/internal/domains"
	"github.com/cyantarek/fizz/internal/pkg/logger"
	"time"
)

type EmailSenders map[string]EmailSender

type EmailService struct {
	emailRepository EmailRepository
	senders         EmailSenders
}

type EmailRepository interface {
	Store(ctx context.Context, email domains.Email) error
	LookupStatus(ctx context.Context, emailID domains.ID) (*domains.Email, error)
	NextEmailID(ctx context.Context) domains.ID
}

type EmailSender interface {
	Name() string
	Send(ctx context.Context, email domains.Email) (domains.ID, error)
}

type EmailStats interface {
	GetStats(ctx context.Context) ([]domains.Stats, error)
}

func (e EmailService) LookupStatus(ctx context.Context, id string) (*LookupEmail, error) {
	emailDomain, err := e.emailRepository.LookupStatus(ctx, domains.NewID(id))
	if err != nil {
		return nil, err
	}

	var out LookupEmail

	out.ID = emailDomain.Id().String()
	out.From = emailDomain.From().Address()
	out.EmailBackend = string(emailDomain.EmailBackend())
	out.Status = emailDomain.Status().String()

	for _, addr := range emailDomain.To() {
		out.To = append(out.To, addr.Address())
	}

	return &out, nil
}

func (e EmailService) Send(ctx context.Context, email SendEmail, backend string) (string, error) {
	// build the domains model from dto
	from := domains.NewEmailAddress(email.From)

	var to []domains.EmailAddress
	for _, v := range email.To {
		to = append(to, domains.NewEmailAddress(v))
	}

	subject := domains.NewSubject(email.Subject)
	body := domains.NewMessageBody(email.Body)

	emailID := e.emailRepository.NextEmailID(ctx)

	emailDomain, err := domains.NewEmail(emailID, from, to, nil, subject, body)
	if err != nil {
		return "", err
	}

	emailDomain.SetEmailBackend(domains.EmailBackend(backend))

	// queue email asynchronously
	go func() {
		emailDomain.MarkAsQueued()

		var refID domains.ID

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

	return emailID.String(), nil
}

func NewEmailService(sender EmailSender, repository EmailRepository) *EmailService {
	return &EmailService{senders: map[string]EmailSender{
		sender.Name(): sender,
	}, emailRepository: repository}
}

type SendEmail struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Body    string
}

type LookupEmail struct {
	ID           string
	From         string
	To           []string
	EmailBackend string
	Status       string
}
