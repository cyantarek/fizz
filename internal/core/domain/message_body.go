package domain

type MessageBody struct {
	body string
}

func (mb MessageBody) Value() string {
	return mb.body
}

func NewMessageBody(body string) MessageBody {
	return MessageBody{body: body}
}
