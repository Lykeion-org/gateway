package handlers

import (
	"net/http"
	"strconv"

	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostAnalyzeText queues a text for analysis. The requester is the authenticated
// user, as set by the RequireAccessLevel middleware.
func (lh *LanguageHandler) PostAnalyzeText(c *gin.Context) {
	var req models.AnalyzeTextRequestClient
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, models.ErrInvalidRequest, err.Error())
		return
	}

	res, err := lh.Client.AnalyzeTextAndAddToSuggestions(c, &lexora.AnalyzeTextAndAddToSuggestionsRequest{
		Text:         req.Text,
		RequesterUid: c.GetString("user_id"),
		Language:     shared.Language_LANGUAGE_EN,
	})
	if err != nil {
		respondWithGrpcError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, models.AnalyzeTextResponse{JobId: res.GetJobUid()})
}

// GetWordSuggestions lists pending word suggestions. Query: page (1-based), pageSize.
func (lh *LanguageHandler) GetWordSuggestions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "50"))

	res, err := lh.Client.ListAllWordSuggestion(c, &lexora.ListAllWordSuggestionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	if err != nil {
		respondWithGrpcError(c, err)
		return
	}

	out := models.WordSuggestionList{
		Suggestions: make([]models.WordSuggestion, 0, len(res.GetSuggestions())),
		TotalCount:  int(res.GetTotalCount()),
	}
	for _, s := range res.GetSuggestions() {
		out.Suggestions = append(out.Suggestions, ProtoToWordSuggestion(s))
	}
	c.JSON(http.StatusOK, out)
}

func respondWithGrpcError(c *gin.Context, err error) {
	st, _ := status.FromError(err)
	switch st.Code() {
	case codes.InvalidArgument:
		respondWithError(c, http.StatusBadRequest, models.ErrInvalidRequest, st.Message())
	case codes.Unavailable:
		respondWithError(c, http.StatusServiceUnavailable, models.ErrInternal, "text analysis is unavailable")
	default:
		respondWithError(c, http.StatusInternalServerError, models.ErrInternal, "internal error")
	}
}
