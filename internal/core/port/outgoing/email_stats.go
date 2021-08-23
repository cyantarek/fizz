package outgoing

import (
	"context"
	"fizz/internal/core/domain"
)

type EmailStats interface {
	GetStats(ctx context.Context) ([]domain.Stats, error)
}
