package mail

import (
	"google.golang.org/api/gmail/v1"
)

type Mail struct {
	From    string
	Date    string
	Subject string
	Body    string
	Images  []Image
}

func NewMail(subject string, msg *gmail.Message) *Mail {
	return &Mail{
		Subject: subject,
		From:    msg.Payload.Headers[0].Value,
		Date:    msg.Payload.Headers[1].Value,
		Body:    msg.Snippet,
	}
}

type Image struct {
	Source      string
	Description string
}
