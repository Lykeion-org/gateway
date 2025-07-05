package controllers

import (
	"github.com/Lykeion-org/gateway/internal/handlers"
	"github.com/Lykeion-org/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AddUserGroups(rg *gin.RouterGroup, h handlers.UserHandler) {
	user := rg.Group("/user")

	user.Use(middleware.RequireAccessLevel(1, h.GetAuthClient()))
	
	user.GET("/", h.GetUserInfo)
}