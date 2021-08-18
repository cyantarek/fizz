package config

import (
	"os"
)

type Config struct {
	HttpPort       string
	EmailBackend   string
	MailgunDomain  string
	MailgunAPIKey  string
	SendgridAPIKey string
	DBHost         string
	DBPort         string
	DBUsername     string
	DBPassword     string
	DBName         string
}

func EnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}
