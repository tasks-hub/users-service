package app

import (
	"fmt"

	"github.com/tasks-hub/users-service/internal/config"
	"github.com/tasks-hub/users-service/internal/handlers"
	"github.com/tasks-hub/users-service/internal/service"
	"github.com/tasks-hub/users-service/internal/store"

	"github.com/gin-gonic/gin"
)

type App struct {
	server      *gin.Engine
	port        string
	userHandler handlers.UserHandler
}

func NewServer(cfg config.Config) (*App, error) {
	r := gin.Default()

	vGroup := r.Group(cfg.UserApiVersion)
	healthHandler := handlers.NewHealthHandler()
	vGroup.GET("/health", healthHandler.Health)

	userStore, err := store.NewUserStoreFactory(cfg)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	vGroup.POST("/users", userHandler.CreateUser)
	vGroup.GET("/users/:id", userHandler.GetUserByID)
	vGroup.POST("/users/authenticate", userHandler.GetUserByEmail)
	vGroup.PUT("/users/:id", userHandler.UpdateUserProfile)
	vGroup.PUT("/users/:id/password", userHandler.ChangePassword)
	vGroup.DELETE("/users/:id", userHandler.DeleteUser)

	return &App{
		server:      r,
		userHandler: userHandler,
		port:        fmt.Sprintf(":%s", cfg.UserServicePort),
	}, nil
}

func (app *App) Run() error {
	return app.server.Run(app.port)
}
