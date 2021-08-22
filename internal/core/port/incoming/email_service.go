package incoming

import (
	"context"
	"fizz/internal/core/application/dto"
)

type EmailService interface {
	Send(ctx context.Context, email dto.SendEmail, backend string) (string, error)
	LookupStatus(ctx context.Context, id string) (*dto.LookupEmail, error)
}
