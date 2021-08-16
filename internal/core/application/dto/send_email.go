package dto

type SendEmail struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Body    string
}
