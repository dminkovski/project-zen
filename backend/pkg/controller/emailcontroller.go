package controller

import (
	"fmt"
	"net/http"
	"project-zen/pkg/auth"
	"project-zen/pkg/mail"

	"github.com/gin-gonic/gin"
)

/*
The EmailController provides the APIs for the frontend to get the summary of all newsletters
*/
type EmailController struct {
	GetEmailsRoute string
	Authenticator  *auth.OAuth
}

func NewEmailController(auth *auth.OAuth) *EmailController {
	return &EmailController{
		GetEmailsRoute: "/emails",
		Authenticator:  auth,
	}
}

func (controller *EmailController) GetEmails(c *gin.Context) {
	fmt.Println("Hello from \"Get Emails\"")
	client := controller.Authenticator.GetClient()
	if client == nil {
		c.JSON(http.StatusUnauthorized, "No token.")
		return
	}
	mails := mail.ReadGmailEmails(client)
	c.JSON(http.StatusOK, mails)
}
