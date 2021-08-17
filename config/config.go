package config

import (
	"os"
)

type Config struct {
	HttpPort       string
	TLSCertFile    string
	TLSKeyFile     string
	EmailBackend   string
	MailgunDomain  string
	MailgunAPIKey  string
	SendgridAPIKey string
	AuthSkipper    map[string]bool
}

func EnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}
