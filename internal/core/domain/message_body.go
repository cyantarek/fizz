package domain

type MessageBody struct {
	body string
}

func NewMessageBody(body string) MessageBody {
	return MessageBody{body: body}
}
