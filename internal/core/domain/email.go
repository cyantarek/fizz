package domain

type Email struct {
	from        EmailAddress
	to          []EmailAddress
	Cc          []EmailAddress
	subject     Subject
	messageBody MessageBody
}

func (e *Email) SetCc(Cc []EmailAddress) {
	e.Cc = Cc
}

func (e *Email) SetSubject(subject Subject) {
	e.subject = subject
}

func (e *Email) SetMessageBody(messageBody MessageBody) {
	e.messageBody = messageBody
}

func NewEmail(from EmailAddress, to []EmailAddress, cc []EmailAddress, subject Subject, messageBody MessageBody) Email {
	if !from.valid() {

	}

	if to == nil {

	}

	if len(to) == 0 {

	}

	return Email{from: from, to: to, Cc: cc, subject: subject, messageBody: messageBody}
}
