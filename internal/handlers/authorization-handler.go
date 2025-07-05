package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"github.com/gin-gonic/gin"
)


type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authHandler struct {
	client kosmos.KosmosServiceClient
}

func NewAuthHandler(orchestrator kosmos.KosmosServiceClient) *authHandler{
	return &authHandler{
		client: orchestrator,
	}
}


func (h *authHandler)Login(c *gin.Context) {
	var req *models.LoginRequest

	err := c.ShouldBind(&req)
	if err != nil {
		slog.Debug("error getting response from kosmos", "error", err)
		respondWithError(c, http.StatusBadRequest, models.ErrInvalidRequest, "invalid request while logging in")
	}

	res, err := h.client.LoginUser(c, &kosmos.LoginUserRequest{Email: req.Username, Password: req.Password})
	if err != nil {
		slog.Debug("error getting response from kosmos", "error", err)
		respondWithError(c, http.StatusUnauthorized, models.ErrUnauthorized, "a server error occured while attempting to login")
	}

	c.JSON(http.StatusOK, &models.TokenResponse{AccessToken: res.AccessToken, RefreshToken: res.RefreshToken})
}

func (h *authHandler)Register(c *gin.Context) {
	var req *models.LoginRequest

	err := c.ShouldBind(&req)
	if err != nil {
		slog.Debug("error getting response from kosmos", "error", err)
		respondWithError(c, http.StatusBadRequest, models.ErrInvalidRequest, "invalid request while registering")
	}

	res, err := h.client.RegisterUser(c, &kosmos.RegisterUserRequest{Email: req.Username, Password: req.Password})
	if err != nil {
		slog.Debug("error getting response from kosmos", "error", err)
		respondWithError(c, http.StatusInternalServerError, models.ErrInternal, "a server error while attempting to create a user")
	}

	c.JSON(http.StatusOK, &models.TokenResponse{AccessToken: res.AccessToken, RefreshToken: res.RefreshToken})
}