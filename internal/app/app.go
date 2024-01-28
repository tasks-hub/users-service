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

func NewServer(cfg config.Config) (*App, error) {
	r := gin.Default()

	v1Group := r.Group("v1")
	healthHandler := handlers.NewHealthHandler()
	v1Group.GET("/health", healthHandler.Health)

	userStore, err := store.NewUserStoreFactory(cfg.DatabaseType, cfg.Database)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	v1Group.POST("/users", userHandler.CreateUser)
	v1Group.GET("/users/:id", userHandler.GetUserByID)
	v1Group.POST("/users/authenticate", userHandler.GetUserByEmail)
	v1Group.PUT("/users/:id", userHandler.UpdateUserProfile)
	v1Group.PUT("/users/:id/password", userHandler.ChangePassword)
	v1Group.DELETE("/users/:id", userHandler.DeleteUser)

	return &App{
		server:      r,
		userHandler: userHandler,
	}, nil
}

func (app *App) Run() error {
	return app.server.Run()
}
