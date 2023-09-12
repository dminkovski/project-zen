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

type OAuth struct {
	clientId     string
	clientSecret string
	config       oauth2.Config
}

func NewOAuth(clientId string, clientSecret string) *OAuth {
	oauth2Config.ClientID = clientId
	oauth2Config.ClientSecret = clientSecret
	config := oauth2Config
	return &OAuth{
		clientId:     clientId,
		clientSecret: clientSecret,
		config:       config,
	}
}

var (
	oauth2Config = oauth2.Config{
		RedirectURL: "https://project-zen.azurewebsites.net/auth/callback",
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

// Retrieve a token, saves the token, then returns the generated client.
func (auth *OAuth) GetClient() *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		fmt.Printf("No token, please start the oauth flow first.\n")
		return nil
	}
	return auth.config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func (auth *OAuth) getTokenFromWeb() *oauth2.Token {
	authURL := auth.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Printf("Unable to read authorization code: %v\n", err)
	}

	tok, err := auth.config.Exchange(context.TODO(), authCode)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web: %v\n", err)
	}
	return tok
}

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

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v\n", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func StartOAuthFlow(c *gin.Context, auth *OAuth) {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		// Create the URL to initiate the OAuth 2.0 flow
		url := auth.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

		// Redirect the user to the Google OAuth consent screen
		c.Redirect(http.StatusFound, url)
	}
	fmt.Printf("Token already saved: %v\n", tok)
	c.JSON(http.StatusOK, tok)
}

func AuthCallback(c *gin.Context, auth *OAuth) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}

	token, err := auth.config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	saveToken("token.json", token)

	c.JSON(http.StatusOK, gin.H{"message": "Authorization successful"})
}
