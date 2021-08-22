package outgoing

import (
	"context"
	"fizz/internal/core/domain"
)

type EmailRepository interface {
	Store(ctx context.Context, email domain.Email) error
	LookupStatus(ctx context.Context, emailID domain.ID) (*domain.Email, error)
	NextEmailID(ctx context.Context) domain.ID
}
