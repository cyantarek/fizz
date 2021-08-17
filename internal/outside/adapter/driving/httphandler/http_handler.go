package httphandler

import (
	"fizz/internal/core/port/incoming"
)

type HTTPHandler struct {
	emailService incoming.EmailService
}

func New(emailService incoming.EmailService) *HTTPHandler {
	return &HTTPHandler{emailService: emailService}
}
