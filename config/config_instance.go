package config

var Cfg = Config{
	HttpPort:       EnvOrDefault("HTTP_PORT", "5000"),
	TLSCertFile:    "",
	TLSKeyFile:     "",
	EmailBackend:   EnvOrDefault("EMAIL_BACKEND", "mailgun"),
	MailgunDomain:  "sandboxc081d8d10815472bad35945be262a094.mailgun.org",
	MailgunAPIKey:  "11800c2a79249b5e8c159d1845c1f44f-1b6eb03d-5e0ff1eb",
	SendgridAPIKey: "",
	AuthSkipper: map[string]bool{
		"/api/v1/auth/register":    true,
		"/auth.Auth/VerifyToken":   true,
		"/product.Product/GetByID": true,
		"/product.Product/Create":  true,
		"/product.Product/List":    true,
		"/api/v1/auth/login":       true,
		"/admin/register":          true,
		"/admin/login":             true,
	},
}
