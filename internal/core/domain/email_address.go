package domain

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) EmailAddress {
	return EmailAddress{address: address}
}

func (e EmailAddress) Address() string {
	return e.address
}

func (e EmailAddress) valid() bool {
	return e.address != ""
}
