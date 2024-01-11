package app

import (
	"gihtub.com/tasks-hub/users-service/internal/config"
	"gihtub.com/tasks-hub/users-service/internal/handlers"
	"gihtub.com/tasks-hub/users-service/internal/services"
	"gihtub.com/tasks-hub/users-service/internal/store"
	"github.com/gin-gonic/gin"
)

type App struct {
	server      *gin.Engine
	userHandler handlers.UserHandler
}

func NewServer(cfg config.Config) *App {
	r := gin.Default()

	healthHandler := handlers.NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	userStore := store.NewInMemoryUserStore()
	userService := services.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)

	return &App{
		server:      r,
		userHandler: userHandler,
	}
}

func (app *App) Run() error {
	return app.server.Run()
}
