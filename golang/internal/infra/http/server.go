package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/madsilver/api-challenge/docs"
	"github.com/madsilver/api-challenge/internal/adapter/handler"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server interface {
	Start()
}

type server struct {
	manager *handler.Manager
}

func NewServer(manager *handler.Manager) Server {
	return &server{manager}
}

// Start @externalDocs.description  OpenAPI
// @title           API Challenge
// @description     Tech Challenge: Performance and Data Analytics.
// @host      localhost:8080
func (s *server) Start() {
	g := gin.Default()
	g.Use(gin.Logger())
	g.POST("users", s.manager.UserHandler.StoreUsers)
	g.GET("superusers", s.manager.UserHandler.GetSuperUsers)
	g.GET("top-countries", s.manager.UserHandler.GetTopCountries)
	g.GET("team-insights", s.manager.UserHandler.GetTeamInsights)
	g.GET("active-users-per-day", s.manager.UserHandler.GetActiveUsers)
	g.GET("evaluation", s.manager.MetricHandler.GetEvaluations)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("Received shutdown signal")
		if err := srv.Close(); err != nil {
			log.Fatalf("Could not close server: %v", err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil {
		msg := "Server closed unexpectedly"
		if !errors.Is(err, http.ErrServerClosed) {
			msg = "Server closed under request"
		}
		log.Println(msg)
	}
	log.Println("Server exiting")
}
