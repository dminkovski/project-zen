package mail

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/api/gmail/v1"
)

func ReadGmailEmails(client *http.Client, extractSummary bool) ([]Mail, error) {
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
		mail, err := extractMailInformation(srv, message, extractSummary)
		if err != nil {
			continue
		}
		mails = append(mails, mail)

	}
	return mails, nil
}

func extractMailInformation(srv *gmail.Service, message *gmail.Message, extractSummary bool) (Mail, error) {
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
	}
	if !extractSummary {
		return Mail{
			Subject: subject,
			Date:    date,
			From:    from,
			Body:    msg.Snippet,
		}, nil
	}

	html, err := getHtml(msg)
	if err != nil {
		return Mail{}, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return Mail{
			Subject: subject,
			Date:    date,
			From:    from,
			Body:    msg.Snippet,
		}, nil
	}
	images := removeDuplicates(getImagesString(doc))
	links := removeDuplicates(getLinksString(doc))

	mail := Mail{
		Subject: subject,
		Date:    date,
		From:    from,
		Body:    msg.Snippet,
		Images:  images,
		Links:   links,
	}
	summary := getMailSummary(mail)
	mail.Summary = summary
	return mail, nil
}

func getHtml(msg *gmail.Message) (string, error) {
	htmlData := msg.Payload.Body.Data
	if htmlData == "" {
		for _, part := range msg.Payload.Parts {
			if part.MimeType == "text/html" {
				htmlData = part.Body.Data
				break
			}
		}
	}
	if htmlData == "" {
		return "", nil
	}
	htmlContent, err := base64.URLEncoding.DecodeString(htmlData)
	if err != nil {
		fmt.Printf("Error decoding HTML content: %v\n", err)
		return "", err
	}

	html := string(htmlContent)
	return html, nil
}

func getImages(doc *goquery.Document) []Image {
	images := make([]Image, 0)
	doc.Find("img").Each(func(index int, imgTag *goquery.Selection) {
		var imageSource string
		var description string
		src, exists := imgTag.Attr("src")
		if exists {
			imageSource = src
		}
		alt, exists := imgTag.Attr("alt")
		if exists {
			description = alt
		}
		if description != "" && imageSource != "" {
			image := Image{
				Source:      imageSource,
				Description: description,
			}
			images = append(images, image)
		}
	})
	return images
}

func getImagesString(doc *goquery.Document) []string {
	images := make([]string, 0)
	doc.Find("img").Each(func(index int, imgTag *goquery.Selection) {
		var description string
		alt, exists := imgTag.Attr("alt")
		if exists {
			description = alt
		}
		if description != "" {
			images = append(images, description)
		}
	})
	return images
}

func getLinks(doc *goquery.Document) []Link {
	var hrefLink string
	links := make([]Link, 0)
	doc.Find("a").Each(func(index int, aTag *goquery.Selection) {
		href, exists := aTag.Attr("href")
		if exists {
			hrefLink = href
		}
		text := aTag.Text()
		link := Link{
			HRef:        hrefLink,
			Description: text,
		}
		links = append(links, link)
	})
	return links
}

func getLinksString(doc *goquery.Document) []string {
	links := make([]string, 0)
	doc.Find("a").Each(func(index int, aTag *goquery.Selection) {
		text := aTag.Text()
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "  ", "", -1)
		if text != "" && !containsOnlyWhitespace(text) {
			links = append(links, text)
		}
	})
	return links
}

func containsOnlyWhitespace(inputStr string) bool {
	// Use a regular expression to check for white spaces
	return regexp.MustCompile(`^\s*$`).MatchString(inputStr)
}

func getMailSummary(mail Mail) string {
	body := SummaryRequest{
		Message: mail.Body,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}
	return getSummary(jsonData)
}

func GetSummaryOfMails(mails []Mail) string {
	allMails := ""
	for _, mail := range mails {
		allMails = allMails + ", " + mail.Body
	}
	body := SummaryRequest{
		Message: allMails,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}
	return getSummary(jsonData)
}

func getSummary(jsonData []byte) string {
	url := "https://projectzen.azurewebsites.net"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("Content-Type", "text/html; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("POST request was successful")
		return string(responseBody)
	} else {
		fmt.Printf("POST request failed with status code %d\n", resp.StatusCode)
	}
	return ""
}

func markMessageAsRead(srv *gmail.Service, userId, messageId string) error {
	modifyRequest := &gmail.ModifyMessageRequest{
		RemoveLabelIds: []string{"UNREAD"},
	}
	_, err := srv.Users.Messages.Modify(userId, messageId, modifyRequest).Do()
	return err
}

func removeDuplicates(input []string) []string {
	// Create a map to store unique elements
	uniqueMap := make(map[string]bool)
	result := []string{}

	// Iterate through the input slice
	for _, item := range input {
		// If the item is not in the map, add it to the result slice
		if _, ok := uniqueMap[item]; !ok {
			result = append(result, item)
			uniqueMap[item] = true
		}
	}

	return result
}

type SummaryRequest struct {
	Message string `json:"message"`
}
