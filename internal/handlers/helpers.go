package handlers

import (
	"errors"
	"strings"
	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/gin-gonic/gin"
)


func respondWithError(c *gin.Context, status int, code models.ErrorCode, message string){
    c.JSON(status, models.ErrorResponse{
        Error: models.ErrorBody{
            Code:    code,
            Message: message,
        },
    })
}


func extractAccessTokenFromHeader(c *gin.Context) (string, error){
	authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return "", errors.New("authorization header missing")
    }

    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
        return "", errors.New("invalid authorization header format")
    }

    return parts[1], nil
}

func extractRefreshTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("'_lykeion_rtoken'")
    if err != nil {
        return "", errors.New("refresh token cookie is missing")
    }
    return token, nil
}

func abortWithError(c *gin.Context, status int, code models.ErrorCode, message string){
    c.AbortWithStatusJSON(status, models.ErrorResponse{
        Error: models.ErrorBody{
            Code:    code,
            Message: message,
        },
    })
}