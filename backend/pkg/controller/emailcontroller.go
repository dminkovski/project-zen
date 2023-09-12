package controller

import (
	"encoding/json"
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
	BlobClient     *auth.StorageClient
}

func NewEmailController(auth *auth.OAuth, blobClient *auth.StorageClient) *EmailController {
	return &EmailController{
		GetEmailsRoute: "/emails",
		Authenticator:  auth,
		BlobClient:     blobClient,
	}
}

func (controller *EmailController) GetEmails(c *gin.Context) {
	fmt.Println("Hello from \"Get Emails\"")
	client := controller.Authenticator.GetClient()
	if client == nil {
		c.JSON(http.StatusUnauthorized, "No token.")
		return
	}
	mails, err := mail.ReadGmailEmails(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(mails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	controller.BlobClient.UploadTextToBlob(jsonData)

	c.JSON(http.StatusOK, mails)
}
