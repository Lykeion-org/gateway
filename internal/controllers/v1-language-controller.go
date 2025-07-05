package controllers

import (
	"github.com/Lykeion-org/gateway/internal/handlers"
	"github.com/Lykeion-org/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)


func AddLanguageRoutes(rg *gin.RouterGroup, h *handlers.LanguageHandler) {
	const requiredEditLevel int32 = 3;
	
	language := rg.Group("/language")


	ref := language.Group("/referent")
	ref.GET("/:uid", h.GetReferent)
	ref.POST("/",middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostReferent)
	ref.PUT("/", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PutReferent)
	ref.GET("/search", h.SearchReferent)

	sym := language.Group("/symbol")
	sym.GET("/:uid", h.GetSymbol)
	sym.POST("/", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostSymbol)
	sym.PUT("/",middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PutSymbol)

	wrd := language.Group("/word")
	wrd.GET("/:uid", h.GetWord)
	wrd.POST("/", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostWord)
	wrd.PUT("/", middleware.RequireAccessLevel(requiredEditLevel, h.Client),h.PutWord)


	language.POST("/link-referent-symbol", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostReferentSymbolLink)
	language.POST("/link-symbol-word", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostSymbolWordLink)
	language.POST("/set-word-as-lemma", middleware.RequireAccessLevel(requiredEditLevel, h.Client), h.PostWordLemmaLink)
}