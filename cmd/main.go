package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/Lykeion-org/gateway/internal/controllers"
	"github.com/Lykeion-org/gateway/internal/grpc"
	"github.com/Lykeion-org/gateway/internal/handlers"
	"github.com/Lykeion-org/go-shared/pkg/config"
	"github.com/Lykeion-org/go-shared/pkg/helpers"
	"github.com/gin-gonic/gin"
)


func main(){

	log.Println("Starting application...\nLoading config file")

	cfg, err := config.LoadConfigFile[Config]("../dev_config.yaml")
	if err != nil {
		slog.Error("Failed to load application configuration file", "error", err)
		panic(err)
	}
	helpers.InitializeStandardLogger("debug")
	
	slog.Info("Opening grpc clients...")
	orchestrator := grpc.NewKosmosClient(cfg.OrchestratorTarget); defer orchestrator.StopClient()
	slog.Info("Opened grpc connection to orchestrator service", "target", cfg.OrchestratorTarget)
	authHandlerV1 := handlers.NewAuthHandler(orchestrator.Client)
	userHandlerV1 := handlers.NewUserHandler(orchestrator.Client)
	languageHandlerV1 := handlers.NewLanguageHandler(orchestrator.Client)


	slog.Info("Initializing REST api...")
	router := gin.Default()

	//Adding routes
	v1 := router.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "Pong"})
	})
	controllers.AddAuthRoutes(v1, authHandlerV1)
	controllers.AddUserGroups(v1, userHandlerV1)
	controllers.AddLanguageRoutes(v1, languageHandlerV1)

	router.Run(cfg.ApiPort)
}