package driven

import (
	"context"
	"fizz/internal/core/domain"
	"fizz/internal/pkg/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type EmailPostgres struct {
	client *sqlx.DB
}

type emailPostgres struct {
	ID           string         `db:"id"`
	ReferenceID  string         `db:"reference_id"`
	From         string         `db:"from_address"`
	To           pq.StringArray `db:"to_address"`
	Cc           pq.StringArray `db:"cc"`
	Subject      string         `db:"subject"`
	Body         string         `db:"body"`
	Status       int            `db:"status"`
	EmailBackend string         `db:"email_backend"`
}

var (
	insertQuery = "INSERT INTO emails (id, reference_id, from_address, to_address, cc, subject, body, status, email_backend) VALUES(:id, :reference_id, :from_address, :to_address, :cc, :subject, :body, :status, :email_backend)"
	selectQuery = "SELECT * FROM emails WHERE id = $1"
)

func (e EmailPostgres) Store(ctx context.Context, email domain.Email) error {
	var to []string

	for _, v := range email.To() {
		to = append(to, v.Address())
	}

	result, err := e.client.NamedExecContext(ctx, insertQuery, emailPostgres{
		ID:           email.Id().String(),
		ReferenceID:  email.ReferenceID().String(),
		From:         email.From().Address(),
		To:           to,
		Subject:      email.Subject().Value(),
		Body:         email.MessageBody().Value(),
		Status:       int(email.Status()),
		EmailBackend: string(email.EmailBackend()),
	})
	if err != nil {
		return err
	}

	logger.Log.Infoln(result.RowsAffected())

	return nil
}

func (e EmailPostgres) LookupStatus(ctx context.Context, emailID domain.ID) (*domain.Email, error) {
	var out emailPostgres

	err := e.client.GetContext(ctx, &out, selectQuery, emailID.String())
	if err != nil {
		return nil, err
	}

	from := domain.NewEmailAddress(out.From)

	var to []domain.EmailAddress
	for _, v := range out.To {
		to = append(to, domain.NewEmailAddress(v))
	}

	subject := domain.NewSubject(out.Subject)
	body := domain.NewMessageBody(out.Body)
	emailBackend := domain.EmailBackend(out.EmailBackend)

	emailDomain, err := domain.NewEmail(domain.NewID(out.ID), from, to, nil, subject, body)
	if err != nil {
		return nil, err
	}

	emailDomain.SetEmailBackend(emailBackend)

	return &emailDomain, nil
}

func (e EmailPostgres) NextEmailID(ctx context.Context) domain.ID {
	return domain.NewID(uuid.NewString())
}

func NewEmailPostgres(client *sqlx.DB) *EmailPostgres {
	return &EmailPostgres{client: client}
}
