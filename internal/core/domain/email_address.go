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
	if e.address == "" {
		return false
	}

	return true
}
