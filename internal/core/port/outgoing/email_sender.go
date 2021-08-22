package outgoing

import (
	"context"
	"fizz/internal/core/domain"
)

type EmailSender interface {
	Name() string
	Send(ctx context.Context, email domain.Email) (domain.ID, error)
}
