package controller

import (
	"fmt"
	"net/http"
	"os"
	"project-zen/pkg/auth"
	"project-zen/pkg/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

/*
The AuthController provides the APIs for the frontend to get the discounts of all newsletters
*/
type AuthController struct {
	AuthCallbackRoute   string
	StartOAuthFlowRoute string
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthCallbackRoute:   "/auth/callback",
		StartOAuthFlowRoute: "/auth/start-outh-flow",
	}
}

func (controller *AuthController) AuthCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}
	c.JSON(http.StatusOK, code)
}

func (controller *AuthController) StartOAuthFlow(c *gin.Context) {
	//auth.StartOAuthFlow(c)
	//ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v\n", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v\n", err)
	}
	client := auth.GetClient(config)
	mail.ReadGmailEmails(client)

	c.JSON(http.StatusOK, "")
}
