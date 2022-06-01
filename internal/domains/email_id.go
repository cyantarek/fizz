package domains

type ID struct {
	id string
}

func (e ID) String() string {
	return e.id
}

func NewID(id string) ID {
	return ID{id: id}
}
