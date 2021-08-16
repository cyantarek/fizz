package outgoing

import (
	"context"
	"fizz/internal/core/domain"
)

type EmailSender interface {
	Send(ctx context.Context, email domain.Email) error
}
