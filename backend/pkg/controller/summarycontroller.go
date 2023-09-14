package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
The SummaryController provides the APIs for the frontend to get the summary of all newsletters
*/
type SummaryController struct {
	GetSummaryRoute string
}

func NewSummaryController() *SummaryController {
	return &SummaryController{
		GetSummaryRoute: "/summary",
	}
}

func (controller *SummaryController) GetSummary(c *gin.Context) {
	fmt.Println("Hello from \"Get Summary\"")
	summary := "This is the great summary"
	c.JSON(http.StatusOK, summary)
}
