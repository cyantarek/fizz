package domain

type Status int

const (
	QUEUED Status = iota
	SENT
	FAILED
	UNKNOWN
)
