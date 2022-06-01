package handlers

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/cyantarek/fizz/internal/services/emails"
	"github.com/cyantarek/fizz/internal/services/marketing"
)

type EmailService interface {
	Send(ctx context.Context, email emails.SendEmail, backend string) (string, error)
	LookupStatus(ctx context.Context, id string) (*emails.LookupEmail, error)
}

type MarketingService interface {
	GetCompleteStats(ctx context.Context) ([]marketing.Stats, error)
}

type Handlers struct {
	router           *echo.Echo
	emailService     EmailService
	marketingService MarketingService
}

func (h Handlers) Start(addr string) error {
	h.routes()

	return h.router.Start(addr)
}

func New(router *echo.Echo, emailService EmailService, marketingService MarketingService) *Handlers {
	return &Handlers{router: router, emailService: emailService, marketingService: marketingService}
}
