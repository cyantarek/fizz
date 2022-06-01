package domains

type Status int

var statusMap = map[Status]string{
	0: "QUEUED",
	1: "SENT",
	2: "FAILED",
	3: "UNKNOWN",
}

func (s Status) String() string {
	return statusMap[s]
}
