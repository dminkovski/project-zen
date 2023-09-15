package aoai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project-zen/pkg/mail"
)

const (
	Aoai = "hi"
)

type AoaiClient struct {
	basePath   string
	httpClient *http.Client
}

func NewAoaiClient() *AoaiClient {
	return &AoaiClient{
		basePath:   "https://projectzen.azurewebsites.net",
		httpClient: &http.Client{},
	}
}

func (aoaiClient *AoaiClient) SummarizeEmail(email *mail.Mail) string {
	/*body := SummarizeEmailRequest{
		From:    mail.From,
		Date:    mail.Date,
		Subject: mail.Subject,
		Body:    mail.Body,
		Links:   mail.Links,
		Images:  mail.Images,
	}*/
	body := SummaryRequest{
		Message: email.Body,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}
	url := aoaiClient.basePath

	return aoaiClient.callAoai(url, jsonData)
}

func (aoaiClient *AoaiClient) SummarizeEmails(emails []*mail.Mail) string {
	if len(emails) == 0 {
		return ""
	}
	allMails := ""
	for _, mail := range emails {
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
	url := aoaiClient.basePath

	return aoaiClient.callAoai(url, jsonData)
}

func (aoaiClient *AoaiClient) callAoai(url string, jsonData []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("Content-Type", "text/html; charset=utf-8")

	resp, err := aoaiClient.httpClient.Do(req)
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

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("POST request failed with status code %d\n", resp.StatusCode)
		return ""
	}
	fmt.Println("POST request was successful")
	return string(responseBody)
}

type SummarizeEmailRequest struct {
	From    string
	Date    string
	Subject string
	Body    string
	Images  []string `json:"Images,omitempty"`
	Links   []string `json:"Links,omitempty"`
}

type SummaryRequest struct {
	Message string `json:"message"`
}
