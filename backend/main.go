package main

import (
	"context"
	"fmt"
	"os"

	"project-zen/pkg/controller"
	"project-zen/pkg/jobs"

	"github.com/gin-gonic/gin"
)

func main() {
	clientId := os.Getenv("clientId")
	fmt.Printf("Client Id: %v\n", clientId)

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

	// Creating the auth controller
	authController := controller.NewAuthController()
	router.GET(authController.AuthCallbackRoute, authController.AuthCallback)
	router.GET(authController.StartOAuthFlowRoute, authController.StartOAuthFlow)

	inboxZenRouter := router.Group(httpPathPrefix)

	// Creating the summaryController for the APIs consumed by the frontend
	summaryController := controller.NewSummaryController()
	discountsController := controller.NewDiscountsController()

	// Setting the routes for the APIs consumed by the frontend
	inboxZenRouter.GET(summaryController.GetSummaryRoute, summaryController.GetSummary)
	inboxZenRouter.GET(discountsController.GetDiscountsRoute, discountsController.GetDiscounts)

	router.Run(port)
}
