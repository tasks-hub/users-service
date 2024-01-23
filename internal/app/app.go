package app

import (
	"github.com/tasks-hub/users-service/internal/config"
	"github.com/tasks-hub/users-service/internal/handlers"
	"github.com/tasks-hub/users-service/internal/service"
	"github.com/tasks-hub/users-service/internal/store"

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
	userService := service.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUserProfile)
	r.PUT("/users/:id/password", userHandler.ChangePassword)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	return &App{
		server:      r,
		userHandler: userHandler,
	}
}

func (app *App) Run() error {
	return app.server.Run()
}
