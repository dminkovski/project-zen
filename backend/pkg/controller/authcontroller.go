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
}
