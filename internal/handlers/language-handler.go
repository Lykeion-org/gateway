package handlers

import (
	"context"
	"net/http"
	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	Client kosmos.KosmosServiceClient
}
func NewLanguageHandler(orchestrator kosmos.KosmosServiceClient) *LanguageHandler {
	return &LanguageHandler{Client: orchestrator}
}

func(lh *LanguageHandler) GetReferent(c *gin.Context){
	resp, err := lh.Client.GetReferent(context.Background(), &lexora.GetReferentRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
	
	ref := ProtoToReferent(resp)
		c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) PostReferent(c *gin.Context){
	var req models.CreateReferentRequest
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &lexora.CreateReferentRequest{			
		EnReference: req.EnReference,
		ImageSource: &req.ImageSource,
	}

	res, err := lh.Client.CreateReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return }

	ref := ProtoToReferent(res)
	c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) PutReferent(c *gin.Context){
	var req models.Referent
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.UpdateReferentRequest{
		Referent: &shared.Referent{
		Uid: req.Uid,
		EnReference: req.EnReference,
		ImageSource: req.ImageSource,
		},
	}

	res, err := lh.Client.UpdateReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToReferent(res)
			
	c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) SearchReferent(c *gin.Context){
	query := c.Query("query")

	resp, err := lh.Client.FindReferents(context.Background(), &lexora.FindReferentsRequest{Query: query, SearchMode: lexora.SearchMode_SEARCH_MODE_WORD})
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	referents := ProtoToReferentList(resp.GetReferents())
	c.JSON(http.StatusOK, referents)

}

func(lh *LanguageHandler) GetSymbol(c *gin.Context){
	res, err := lh.Client.GetSymbol(context.Background(), &lexora.GetSymbolRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	

	sym := ProtoToSymbol(res)
		c.JSON(http.StatusOK, sym)
}

func(lh *LanguageHandler) PostSymbol(c *gin.Context){
	var req models.CreateSymbolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &lexora.CreateSymbolRequest{
		ReferentUid: req.ReferentUid,			
		Language:    shared.Language(req.Language),
		SymbolType:  shared.SymbolType(req.SymbolType),
	}

	res, err := lh.Client.CreateSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return 
	}

	sym := ProtoToSymbol(res)
	c.JSON(http.StatusOK, sym)

}

func(lh *LanguageHandler) PutSymbol(c *gin.Context){
	var req models.Symbol
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.UpdateSymbolRequest{
		Symbol: &shared.Symbol{
		Uid: req.Uid,
		Language: shared.Language(req.Language),
		SymbolType: shared.SymbolType(req.SymbolType),
		},
	}
	res, err := lh.Client.UpdateSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToSymbol(res)
			
	c.JSON(http.StatusOK, ref)
}


func(lh *LanguageHandler) GetWord(c *gin.Context){
	res, err := lh.Client.GetWord(context.Background(), &lexora.GetWordRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		wrd := ProtoToWord(res)
		c.JSON(http.StatusOK, wrd)
}

func(lh *LanguageHandler) PostWord(c *gin.Context){
	var req models.CreateWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &lexora.CreateWordRequest{
		SymbolUid: 		req.SymbolUid,
		Word: req.Word,
		WordType: shared.WordType(req.WordType),
		SoundSource: &req.SoundSource,
		Ipa: &req.Ipa,
	}

	res, err := lh.Client.CreateWord(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return 
	}

	wrd := ProtoToWord(res)
	c.JSON(http.StatusOK, wrd)
}

func(lh *LanguageHandler) PutWord(c *gin.Context){
	var req models.Word
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.UpdateWordRequest{
		Word: &shared.Word{
		Uid: req.Uid,
		Word: req.Word,
		SoundSource: req.SoundSource,
		Ipa: req.Ipa,
		},
	}

	res, err := lh.Client.UpdateWord(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToWord(res)
			
	c.JSON(http.StatusOK, ref)
}


func(lh *LanguageHandler) PostReferentSymbolLink(c *gin.Context){
	var req models.LinkSymbolToReferentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.LinkSymbolToReferentRequest{
		ReferentUid: req.ReferentUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.LinkSymbolToReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

func(lh *LanguageHandler) PostSymbolWordLink(c *gin.Context){
	var req models.LinkWordToSymbolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.LinkWordToSymbolRequest{
		WordUid: req.WordUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.LinkWordToSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

func(lh *LanguageHandler) PostWordLemmaLink(c *gin.Context){
	var req models.SetWordAsLemmaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &lexora.SetSymbolLemmaRequest{
		WordUid: req.WordUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.SetSymbolLemma(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

