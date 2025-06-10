package api

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	pb "github.com/Lykeion/gateway/internal/grpc/generated"
)

type api struct {
	engine *gin.Engine
	port string
	grpcClient pb.LanguageServiceClient
}

type CreateReferentRequest struct {
	EnReference string `json:"enReference" binding:"required"`
	ImageSource string `json:"imageSource"`
}

type CreateSymbolRequest struct {
	ReferentUid string `json:"referentUid"`
    Language int `json:"language"`
    SymbolType int `json:"symbolType"`
}

type LinkSymbolToReferentRequest struct {
	SymbolUid string `json:"symbolUid"`
	ReferentUid string `json:"referentUid"`
}

type Referent struct {
	Uid string `json:"uid"`
	EnReference string `json:"enReference"`
	ImageSource string `json:"imageSource"`
}



func NewApi(grpcConn pb.LanguageServiceClient) *api {
	r := gin.Default()

	return &api {
		engine: r,
		port: ":8080",
		grpcClient: grpcConn,
	}
}

func (a *api) InitializeApi() {
	a.engine.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "Hello": "World",
        })
    })

	a.engine.GET("/referent/:uid", func(c *gin.Context) {
		resp, err := a.grpcClient.GetReferent(context.Background(), &pb.GetReferentRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
		c.JSON(http.StatusOK, resp)
	})

	a.engine.GET("/symbol/:uid", func(c *gin.Context) {
		resp, err := a.grpcClient.GetSymbol(context.Background(), &pb.GetSymbolRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
		c.JSON(http.StatusOK, resp)
	})

	a.engine.GET("/word/:uid", func(c *gin.Context) {
		resp, err := a.grpcClient.GetWord(context.Background(), &pb.GetWordRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
		c.JSON(http.StatusOK, resp)
	})

	a.engine.POST("/referent", func(c *gin.Context) {
		var req CreateReferentRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		grpcReq := &pb.CreateReferentRequest{
			EnReference: req.EnReference,
			ImageSource: &req.ImageSource,
		}

		res, err := a.grpcClient.CreateReferent(context.Background(), grpcReq); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return }
			
		c.JSON(http.StatusOK, res)
	})

	a.engine.PUT("/referent", func(c *gin.Context) {
		var req Referent
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		grpcReq := &pb.UpdateReferentRequest{
			Referent: &pb.Referent{
				Uid: req.Uid,
				EnReference: req.EnReference,
				ImageSource: req.ImageSource,
			},
		}

		res, err := a.grpcClient.UpdateReferent(context.Background(), grpcReq); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return }
			
		c.JSON(http.StatusOK, res)
	})

	a.engine.POST("/symbol", func(c *gin.Context) {
		var req CreateSymbolRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		grpcReq := &pb.CreateSymbolRequest{
			ReferentUid: req.ReferentUid,
			Language:    pb.Language(req.Language),
			SymbolType:  pb.SymbolType(req.SymbolType),
		}

		res, err := a.grpcClient.CreateSymbol(context.Background(), grpcReq); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return }
			
		c.JSON(http.StatusOK, res)

	})

	a.engine.POST("/link-referent-symbol", func(c *gin.Context) {
		var req LinkSymbolToReferentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		grpcReq := &pb.LinkSymbolToReferentRequest{
			ReferentUid: req.ReferentUid,
			SymbolUid: req.SymbolUid,
		}

		res, err := a.grpcClient.LinkSymbolToReferent(context.Background(), grpcReq); if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return }
			
		c.JSON(http.StatusOK, res)
	})

	a.engine.Run(a.port)
}

