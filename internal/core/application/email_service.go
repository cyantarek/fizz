package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/port/outgoing"
)

type EmailService struct {
	sender outgoing.EmailSender
}

func (e EmailService) Send(ctx context.Context, email dto.SendEmail) error {
	panic("implement me")
}

func NewEmailService(sender outgoing.EmailSender) *EmailService {
	return &EmailService{sender: sender}
}
