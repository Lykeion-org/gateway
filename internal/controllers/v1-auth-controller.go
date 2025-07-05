package controllers

import (
	"github.com/Lykeion-org/gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(rg *gin.RouterGroup, h handlers.AuthHandler) {
	auth := rg.Group("/auth")
	
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}