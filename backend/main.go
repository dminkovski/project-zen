package main

import (
	"context"
	"fmt"
	"os"

	"project-zen/pkg/auth"
	"project-zen/pkg/controller"
	"project-zen/pkg/jobs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	clientId, exists := os.LookupEnv("clientId")
	if !exists {
		fmt.Println("failed to get clientId")
		os.Exit(1)
	}
	fmt.Printf("Client Id: %v\n", clientId)

	clientSecret, exists := os.LookupEnv("clientSecret")
	if !exists {
		fmt.Println("failed to get clientSecret")
		os.Exit(1)
	}

	oauth := auth.NewOAuth(clientId, clientSecret)

	backgroundContext := context.Background()

	// setting up the scheduler
	job := jobs.NewSummarizeNewlettersJob()
	err := job.StartScheduler(backgroundContext)
	if err != nil {
		fmt.Println("failed to start scheduler")
		os.Exit(1)
	}

	// setting up the server
	httpPathPrefix := "/project-zen"
	port := ":8080"
	router := gin.Default()
	router.Use(cors.Default())

	// Creating the auth controller
	authController := controller.NewAuthController(oauth)
	router.GET(authController.AuthCallbackRoute, authController.AuthCallback)
	router.GET(authController.StartOAuthFlowRoute, authController.StartOAuthFlow)

	inboxZenRouter := router.Group(httpPathPrefix)

	// Creating the controller for the APIs consumed by the frontend
	emailController := controller.NewEmailController(oauth, nil)

	// Setting the routes for the APIs consumed by the frontend
	inboxZenRouter.GET(emailController.GetEmailsRoute, emailController.GetEmails)
	inboxZenRouter.GET(emailController.GetEmailsWithSummaryRoute, emailController.GetEmailsWithSummary)

	router.Run(port)
}
