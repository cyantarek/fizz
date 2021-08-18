package domain

type Email struct {
	id           EmailID
	referenceID  EmailID
	from         EmailAddress
	to           []EmailAddress
	cc           []EmailAddress
	subject      Subject
	messageBody  MessageBody
	status       Status
	emailBackend EmailBackend
}

func (e *Email) Id() EmailID {
	return e.id
}

func (e *Email) ReferenceID() EmailID {
	return e.referenceID
}

func (e *Email) Status() Status {
	return e.status
}

func (e *Email) EmailBackend() EmailBackend {
	return e.emailBackend
}

func (e *Email) SetEmailBackend(emailBackend EmailBackend) {
	e.emailBackend = emailBackend
}

func (e *Email) SetReferenceID(referenceID EmailID) {
	e.referenceID = referenceID
}

func (e *Email) MarkAsQueued() {
	e.status = QUEUED
}

func (e *Email) From() EmailAddress {
	return e.from
}

func (e *Email) To() []EmailAddress {
	return e.to
}

func (e *Email) Cc() []EmailAddress {
	return e.cc
}

func (e *Email) Subject() Subject {
	return e.subject
}

func (e *Email) MessageBody() MessageBody {
	return e.messageBody
}

func (e *Email) SetCc(Cc []EmailAddress) {
	e.cc = Cc
}

func (e *Email) SetSubject(subject Subject) {
	e.subject = subject
}

func (e *Email) SetMessageBody(messageBody MessageBody) {
	e.messageBody = messageBody
}

func NewEmail(id EmailID, from EmailAddress, to []EmailAddress, cc []EmailAddress, subject Subject, messageBody MessageBody) (Email, error) {
	if !from.valid() {
		return Email{}, errInvalidFrom
	}

	if to == nil {
		return Email{}, errInvalidTo
	}

	if len(to) == 0 {
		return Email{}, errInvalidTo
	}

	return Email{id: id, from: from, to: to, cc: cc, subject: subject, messageBody: messageBody}, nil
}
