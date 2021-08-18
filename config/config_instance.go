package config

var Cfg = Config{
	HttpPort:       EnvOrDefault("HTTP_PORT", "5000"),
	EmailBackend:   EnvOrDefault("EMAIL_BACKEND", "mailgun"),
	MailgunDomain:  EnvOrDefault("MAILGUN_DOMAIN", ""),
	MailgunAPIKey:  EnvOrDefault("MAILGUN_API_KEY", ""),
	DBHost:         EnvOrDefault("DB_HOST", "localhost"),
	DBPort:         EnvOrDefault("DB_PORT", "5432"),
	DBUsername:     EnvOrDefault("DB_USERNAME", "fizz"),
	DBPassword:     EnvOrDefault("DB_PASSWORD", ""),
	DBName:         EnvOrDefault("DB_NAME", "fizz"),
}
