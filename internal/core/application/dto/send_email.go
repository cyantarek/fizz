package dto

type SendEmail struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Body    string
}

type LookupEmail struct {
	ID           string
	From         string
	To           []string
	EmailBackend string
	Status       string
}
