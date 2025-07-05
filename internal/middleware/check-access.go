package middleware

import (
	"log/slog"
	"net/http"

	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"github.com/gin-gonic/gin"
)

func RequireAccessLevel(requiredLevel int32, svc kosmos.KosmosServiceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := extractAccessTokenFromHeader(c)
		if err != nil {
			slog.Debug("error extracting access token from header", "error", err)
			abortWithError(c, http.StatusUnauthorized, models.ErrUnauthorized, "missing or invalid access token")
			return
		}

		claims, err := svc.GetUserClaims(c, &kosmos.RequestToken{AccessToken: token})
		if err != nil {
			slog.Debug("error checking claims from kosmos service", "error", err)
			abortWithError(c, http.StatusUnauthorized, models.ErrUnauthorized, "could not validate access token")
			return
		}
		c.Set("user_id", claims.UserId)
		c.Set("user_role", claims.Role)

		if claims.Role < requiredLevel {
			slog.Warn(
				"user requested access to forbidden resource", 
				"user",claims.UserId, 
				"requiredLvl", requiredLevel, 
				"actualLvl", claims.Role,
			)
			abortWithError(c, http.StatusForbidden, models.ErrForbidden, "insufficient access level")
			return
		}

		c.Next()
    }
}