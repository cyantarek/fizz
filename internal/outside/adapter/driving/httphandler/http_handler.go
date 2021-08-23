package httphandler

import (
	"fizz/internal/core/port/incoming"
)

type HTTPHandler struct {
	emailService     incoming.EmailService
	marketingService incoming.MarketingService
}

func New(emailService incoming.EmailService, marketingService incoming.MarketingService) *HTTPHandler {
	return &HTTPHandler{emailService: emailService, marketingService: marketingService}
}
