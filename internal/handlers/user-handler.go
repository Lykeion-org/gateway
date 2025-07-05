package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"github.com/gin-gonic/gin"
)
type UserHandler interface {
	GetUserInfo(c *gin.Context)
	GetAuthClient() (kosmos.KosmosServiceClient)
}

type userHandler struct {
	client kosmos.KosmosServiceClient
}
func NewUserHandler(orchestrator kosmos.KosmosServiceClient) UserHandler {
	return &userHandler{client: orchestrator}
}
func (h *userHandler) GetUserInfo(c *gin.Context) {
	uidValue, exists := c.Get("user_id")
	slog.Debug("Get user infor request handler, context uid", "uid", uidValue )
	if !exists {
		abortWithError(c, http.StatusUnauthorized, models.ErrUnauthorized, "user ID missing in context")
		return
	}

	uid, ok := uidValue.(string)
	if !ok {
		abortWithError(c, http.StatusInternalServerError, models.ErrInternal, "user ID in context is not a string")
		return
	}

	res, err := h.client.GetUserInfo(c, &kosmos.GetUserInfoRequest{UserUid: uid})
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, models.ErrInternal, "could not get user info")
		return// Handle error (e.g., log and return 500)
	}

	c.JSON(http.StatusOK, &models.User{
		Uid					: res.GetUserId(),
		EmailAddress		: res.GetEmailAddress(),
		AvatarSource		: res.GetAvatarSource(),
		NativeLanguage		: res.GetNativeLanguage(),
		ReferenceLanguage	: res.GetReferenceLanguage(),
		TargetLanguage		: res.GetTargetLanguage(),
		FirstName			: res.GetFirstName(),
		LastName			: res.GetLastName(),
		ExtraName			: res.GetExtraName(),
		DisplayName			: res.GetDisplayName(),
		Role				: res.GetRole(),
	})
}

func(h *userHandler)GetAuthClient() kosmos.KosmosServiceClient {
	return h.client
}
