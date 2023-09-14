package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
The DiscountsController provides the APIs for the frontend to get the discounts of all newsletters
*/
type DiscountsController struct {
	GetDiscountsRoute string
}

func NewDiscountsController() *DiscountsController {
	return &DiscountsController{
		GetDiscountsRoute: "/discounts",
	}
}

func (controller *DiscountsController) GetDiscounts(c *gin.Context) {
	fmt.Println("Hello from \"Get Discounts\"")
	discounts := "These are your awesome discounts"
	c.JSON(http.StatusOK, discounts)
}
