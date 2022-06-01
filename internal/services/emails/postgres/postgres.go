package postgres

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/cyantarek/fizz/internal/domains"
	"github.com/cyantarek/fizz/internal/pkg/logger"
)

type postgresDB struct {
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

func (e postgresDB) Store(ctx context.Context, email domains.Email) error {
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

func (e postgresDB) LookupStatus(ctx context.Context, emailID domains.ID) (*domains.Email, error) {
	var out emailPostgres

	err := e.client.GetContext(ctx, &out, selectQuery, emailID.String())
	if err != nil {
		return nil, err
	}

	from := domains.NewEmailAddress(out.From)

	var to []domains.EmailAddress
	for _, v := range out.To {
		to = append(to, domains.NewEmailAddress(v))
	}

	subject := domains.NewSubject(out.Subject)
	body := domains.NewMessageBody(out.Body)
	emailBackend := domains.EmailBackend(out.EmailBackend)

	emailDomain, err := domains.NewEmail(domains.NewID(out.ID), from, to, nil, subject, body)
	if err != nil {
		return nil, err
	}

	emailDomain.SetEmailBackend(emailBackend)

	return &emailDomain, nil
}

func (e postgresDB) NextEmailID(ctx context.Context) domains.ID {
	uid := strings.Replace(uuid.NewString(), "-", "", -1)
	return domains.NewID(uid)
}

func New(client *sqlx.DB) *postgresDB {
	return &postgresDB{client: client}
}
