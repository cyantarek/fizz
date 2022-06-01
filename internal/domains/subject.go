package domains

type Subject struct {
	subject string
}

func (s Subject) Value() string {
	return s.subject
}

func NewSubject(subject string) Subject {
	return Subject{subject: subject}
}
