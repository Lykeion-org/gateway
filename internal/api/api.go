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

	a.engine.Run(a.port)
}

