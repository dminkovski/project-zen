package controller

import (
	"project-zen/pkg/auth"

	"github.com/gin-gonic/gin"
)

/*
The AuthController provides the APIs for the frontend to get the discounts of all newsletters
*/
type AuthController struct {
	AuthCallbackRoute   string
	StartOAuthFlowRoute string
	Authenticator       *auth.OAuth
}

func NewAuthController(authenticator *auth.OAuth) *AuthController {
	return &AuthController{
		AuthCallbackRoute:   "/auth/callback",
		StartOAuthFlowRoute: "/auth/start-oauth-flow",
		Authenticator:       authenticator,
	}
}

func (controller *AuthController) AuthCallback(c *gin.Context) {
	auth.AuthCallback(c, controller.Authenticator)
}

func (controller *AuthController) StartOAuthFlow(c *gin.Context) {
	auth.StartOAuthFlow(c, controller.Authenticator)
	//ctx := context.Background()
	/*b, err := os.ReadFile("credentials.json")
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

	c.JSON(http.StatusOK, "")*/
}
