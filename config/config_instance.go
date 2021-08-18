package config

var Cfg = Config{
	HttpPort:       EnvOrDefault("HTTP_PORT", "5000"),
	TLSCertFile:    "",
	TLSKeyFile:     "",
	EmailBackend:   EnvOrDefault("EMAIL_BACKEND", "mailgun"),
	MailgunDomain:  EnvOrDefault("MAILGUN_DOMAIN", ""),
	MailgunAPIKey:  EnvOrDefault("MAILGUN_API_KEY", ""),
	SendgridAPIKey: "",
}
