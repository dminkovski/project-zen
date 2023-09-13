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
	connectionString, exists := os.LookupEnv("storageaccountkey")
	if !exists {
		fmt.Println("failed to get connectionString")
		os.Exit(1)
	}
	blobClient := auth.NewStorageClient(connectionString)

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

	// Creating the summaryController for the APIs consumed by the frontend
	summaryController := controller.NewSummaryController()
	discountsController := controller.NewDiscountsController()
	emailController := controller.NewEmailController(oauth, blobClient)

	// Setting the routes for the APIs consumed by the frontend
	inboxZenRouter.GET(summaryController.GetSummaryRoute, summaryController.GetSummary)
	inboxZenRouter.GET(discountsController.GetDiscountsRoute, discountsController.GetDiscounts)
	inboxZenRouter.GET(emailController.GetEmailsRoute, emailController.GetEmails)

	router.Run(port)
}
