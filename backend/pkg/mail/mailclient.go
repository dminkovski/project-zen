package mail

import (
	"fmt"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

func ReadGmailEmails(client *http.Client) ([]Mail, error) {
	mails := make([]Mail, 0)
	// Create a Gmail API service instance.
	srv, err := gmail.New(client)
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail messages: %v\n", err)
		return mails, err
	}

	/*srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}*/

	// List Gmail messages.
	messages, err := srv.Users.Messages.List("me").Q("is:unread").Do()
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail messages: %v\n", err)
		return mails, err
	}

	fmt.Println("Unread Messages:")
	for _, message := range messages.Messages {
		msg, err := srv.Users.Messages.Get("me", message.Id).Do()
		if err != nil {
			fmt.Printf("Unable to retrieve message details: %v\n", err)
		}

		// Extract the subject from the message headers.
		subject := ""
		from := ""
		date := ""
		for _, header := range msg.Payload.Headers {
			if header.Name == "Subject" {
				subject = header.Value
			}

			if header.Name == "From" {
				from = header.Value
			}

			if header.Name == "Date" {
				date = header.Value
			}

			if header.Name == "Subject" {
				subject = header.Value
			}
		}

		mail := Mail{
			Subject: subject,
			Date:    date,
			From:    from,
			Body:    msg.Snippet,
		}
		mails = append(mails, mail)
		err = markMessageAsRead(srv, "me", message.Id)
		if err != nil {
			fmt.Printf("Unable to mark message as read: %v\n", err)
		} else {
			fmt.Printf("Marked message as read: %s\n", message.Id)
		}
	}
	return mails, nil
}

func markMessageAsRead(srv *gmail.Service, userId, messageId string) error {
	modifyRequest := &gmail.ModifyMessageRequest{
		RemoveLabelIds: []string{"UNREAD"},
	}
	_, err := srv.Users.Messages.Modify(userId, messageId, modifyRequest).Do()
	return err
}
