package domain

type Subject struct {
	subject string
}

func NewSubject(subject string) Subject {
	return Subject{subject: subject}
}
