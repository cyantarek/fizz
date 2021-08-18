package domain

type EmailBackend int

const (
	MAILGUN EmailBackend = iota
	SENDGRID
)
