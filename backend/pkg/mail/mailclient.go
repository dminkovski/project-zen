package mail

import (
	"fmt"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

func ReadGmailEmails(client *http.Client) string {
	// Create a Gmail API service instance.
	srv, err := gmail.New(client)
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail messages: %v\n", err)
	}

	/*srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}*/

	// List Gmail messages.
	messages, err := srv.Users.Messages.List("me").Q("is:unread").Do()
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail messages: %v\n", err)
	}

	body := ""

	fmt.Println("Unread Messages:")
	for _, message := range messages.Messages {
		msg, err := srv.Users.Messages.Get("me", message.Id).Do()
		if err != nil {
			fmt.Printf("Unable to retrieve message details: %v\n", err)
		}

		// Extract the subject from the message headers.
		subject := ""
		for _, header := range msg.Payload.Headers {
			if header.Name == "Subject" {
				subject = header.Value
				break
			}
		}

		fmt.Printf("- Subject: %s\n", subject)
		fmt.Printf("  From: %s\n", msg.Payload.Headers[0].Value) // Assuming the first header is "From"
		fmt.Printf("  Date: %s\n", msg.Payload.Headers[1].Value) // Assuming the second header is "Date"
		fmt.Printf("  Body: %s\n", msg.Snippet)
		body = msg.Snippet
	}
	return body
}
