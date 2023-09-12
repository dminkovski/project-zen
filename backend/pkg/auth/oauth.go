package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	fmt.Println("Getting client")
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Printf("Unable to read authorization code: %v\n", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web: %v\n", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v\n", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

var (
	oauth2Config = oauth2.Config{
		ClientID:     "..",
		ClientSecret: "..",
		RedirectURL:  "https://project-zen.azurewebsites.net/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/gmail.modify",
		},
	}
)

func StartOAuthFlow(c *gin.Context) {
	// Create the URL to initiate the OAuth 2.0 flow
	url := oauth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Redirect the user to the Google OAuth consent screen
	c.Redirect(http.StatusFound, url)
}

func AuthCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		// Handle the case where there's no code
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}

	token, err := oauth2Config.Exchange(c, code)
	if err != nil {
		// Handle the error (e.g., display an error page)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Store the token securely (e.g., in a database) for later use
	saveToken("token.json", token)

	// Redirect the user to a success page or perform further actions
	c.JSON(http.StatusOK, gin.H{"message": "Authorization successful"})
}
