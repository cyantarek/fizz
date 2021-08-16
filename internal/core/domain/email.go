package domain

type Email struct {
	From        EmailAddress
	To          []EmailAddress
	MessageBody MessageBody
}
