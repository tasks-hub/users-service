package app

import (
	"gihtub.com/fabulias/task-management-backend/users-service/internal/config"
	"gihtub.com/fabulias/task-management-backend/users-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *gin.Engine
}

func NewServer(cfg config.Config) *App {
	r := gin.Default()

	healthHandler := handlers.NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	return &App{
		server: r,
	}
}

func (app *App) Run() error {
	return app.server.Run()
}
