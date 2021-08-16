package config

var Cfg = Config{
	HttpPort:         EnvOrDefault("HTTP_PORT", "5000"),
	APIPort:          EnvOrDefault("API_PORT", "5001"),
	GRPCPort:         EnvOrDefault("GRPC_PORT", "5002"),
	DBHost:           EnvOrDefault("DB_PORT", "localhost"),
	DBPort:           EnvOrDefault("DB_PORT", "9432"),
	DBName:           EnvOrDefault("DB_NAME", "boilerplate"),
	DBUser:           EnvOrDefault("DB_USER", "listmonk"),
	DBPassword:       EnvOrDefault("DB_PASSWORD", "listmonk"),
	JWTSecret:        EnvOrDefault("JWT_SECRET", "*234230KJHDS"),
	JWTRefreshSecret: EnvOrDefault("JWT_REFRESH_SECRET", "WE13123??__<>"),
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
	GCPServiceAccount: `{
		"type": "service_account",
		"project_id": "go-standard-boilerplate",
		"private_key_id": "03e9e2a279274a5d5ff61ff007652d46cce83ae1",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDSzcRDZ1WhoQYT\nPg4o7+bfZzi/UqfXI+tGeD2KANTBDjdvYP3MIj/8WBy4W9kR5PW29ZJJ7W/g3N+f\numr0BJ5/PZSZJPhmfQjs2BXwzEGJGq8KW97D6GWPaXWWbTUHYngujQZuudjIfNlR\nNEXgSIVBQ0zQSSClC94/8+HjFwQ/NfZ3Z31vKBXLAYL/bFATj8vhoKz8wA3xlnHf\n3coisBmH+gcye6fuod7hzKgLbi5xzOrAKeOmm/QaRt2gYqFXjug5NYwSU7KWvAPr\nFJCChc+OAXmy1RJnx8YT+2GvkgfYi5dZep6z5kRCZA1OiV1jmN+7rxueDjzlAuZa\nTU2RCxaVAgMBAAECggEAApjqDT4rBVpHU/sLXoXBifRJ0xfYlvsso3R6ZFcYmo+Q\neZ0IlATbHDIh/Ro7USrNq7jYSXzEvB8+oRXWBEsoNjLw8YOohKDyqenyQGu7l6B6\nu8Smhyoe5r10GMzV27hnASGBX23yhEn/UexuDelCMLxwMG42OxRttcK4qaZ4f3Dt\nBmbh8scQMt0CVoaKameIhH1imy6ZUktAB+O/Me+BQb+mhfVdpMQ8A6SFiwaNTWF7\nJteiYuwrimIK3snanPF4whV+rUe7iHDnZlQGuQX3a54SoJxIL/NwhpKHI7C8j3PH\nCVpjAeyGwJz/0q3VdSDPjrg1P1GrkWLHTXhATD9EoQKBgQD8zBb4wg14P+oNeFJF\nmrgjQGutLlzVacvjaN6X3gQfYB6UDP0bZndBn/YE+5Kgu3bWPIv8doMejgg4/wFc\nYvBxLyYDh2drY1wiPB2W3s156mBYZbsP8PhKy+o9/5PRdJ05ezp/GTu4Ukkx22f/\nRMuLS9a0iV19EIG5flVSWzDm8wKBgQDVeXo2/3+dGJ4IVELlRhJgLMtVxyIPGCL/\nhRlCI4ooOrbfvK7nvAO//F75cFQgjRP326qvj5++cZpUMB9fb5n1IoEt1GgNYlIa\n87+z9+TdcPZdho3vNF8+WIaIjtC5uc/WeT+PJIJTFUlveUXlUmdgCjko4zVGWOpA\n00YrCXt+VwKBgFoLmGMjPAkJOyVxJl0Er3JfD/uv+AFMngNy51bRDkbexgWWWtHX\nvcLyZ39+3MvD+qB1EKfszuejT+p9as6tr1eho7i+Emh+C+Nl5mRHGInEomzaT7dt\n7gM2f0l2MAD7uMUwz7VWF5+gL9JChNi4eIg1i1TjWRmjN5ILHfkn4lNhAoGAd14T\n7sjzbL6lL3ceaOHDyK+Di5VsABC5ETQ1qXwB9vjN5VG2Y7IITsQpv0UGZaU9cy0L\nxluDMZSIfWbjRBQ9fIssvJm3DTTbQGLn4RROj1xpBmcE2qroTw5lO9rP7+pswzrN\nRZdeoGKYy5J+ePqY/2T4DVrvIUT6yADbjZkhysUCgYEA1Uh/mduOWb6q61tz8wKn\nE0aHZ4M7Q0HMCHuNbT9wbh/o+dphPaNg8cXd6BfMyIq3rkVDe1GYsLarFyz/tLeO\njuo5/6v/qTtXr7k+ltGe1QUA0NbrCsWZyvBMOEKDAl+v/P50QpIWGOT4TSWcnIX4\nafz7ucMGdlMO1OuMnKM2JyA=\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-37moo@go-standard-boilerplate.iam.gserviceaccount.com",
		"client_id": "110552334845065938152",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-37moo%40go-standard-boilerplate.iam.gserviceaccount.com"
	}`,
	GraphQLPort: EnvOrDefault("HTTP_PORT", "5003"),
}
