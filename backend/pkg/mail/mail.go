package mail

import "google.golang.org/api/gmail/v1"

type Mail struct {
	From    string
	Date    string
	Subject string
	Body    string
}

func NewMail(subject string, msg *gmail.Message) *Mail {
	return &Mail{
		Subject: subject,
		From:    msg.Payload.Headers[0].Value,
		Date:    msg.Payload.Headers[1].Value,
		Body:    msg.Snippet,
	}
}
