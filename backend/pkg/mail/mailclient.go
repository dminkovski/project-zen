package mail

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	// List Gmail messages.
	messages, err := srv.Users.Messages.List("me").Q("is:unread").Do()
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail messages: %v\n", err)
		return mails, err
	}

	fmt.Println("Unread Messages:")
	for _, message := range messages.Messages {
		mail, err := extractMailInformation(srv, message)
		if err != nil {
			continue
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

func extractMailInformation(srv *gmail.Service, message *gmail.Message) (Mail, error) {
	msg, err := srv.Users.Messages.Get("me", message.Id).Do()
	if err != nil {
		fmt.Printf("Unable to retrieve message details: %v\n", err)
		return Mail{}, nil
	}

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

	html, err := getHtml(msg)
	if err != nil {
		return Mail{}, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return Mail{}, nil
	}
	images := getImages(doc)

	mail := Mail{
		Subject: subject,
		Date:    date,
		From:    from,
		Body:    msg.Snippet,
		Images:  images,
	}
	return mail, nil
}

func getHtml(msg *gmail.Message) (string, error) {
	htmlData := msg.Payload.Body.Data
	htmlContent, err := base64.URLEncoding.DecodeString(htmlData)
	if err != nil {
		fmt.Printf("Error decoding HTML content: %v\n", err)
		return "", err
	}

	html := string(htmlContent)
	return html, nil
}

func getImages(doc *goquery.Document) []Image {
	var imageSource string
	var description string
	images := make([]Image, 0)
	// Find and process <img> tags
	doc.Find("img").Each(func(index int, imgTag *goquery.Selection) {
		src, exists := imgTag.Attr("src")
		if exists {
			imageSource = src
		}
		alt, exists := imgTag.Attr("alt")
		if exists {
			description = alt
		}
		image := Image{
			Source:      imageSource,
			Description: description,
		}
		images = append(images, image)
		// Download images?
	})
	return images
}

func markMessageAsRead(srv *gmail.Service, userId, messageId string) error {
	modifyRequest := &gmail.ModifyMessageRequest{
		RemoveLabelIds: []string{"UNREAD"},
	}
	_, err := srv.Users.Messages.Modify(userId, messageId, modifyRequest).Do()
	return err
}
