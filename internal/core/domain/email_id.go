package domain

type EmailID struct {
	id string
}

func (e EmailID) String() string {
	return e.id
}

func NewEmailID(id string) EmailID {
	return EmailID{id: id}
}
