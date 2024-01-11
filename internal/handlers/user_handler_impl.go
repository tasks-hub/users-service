package handlers

import (
	"gihtub.com/tasks-hub/users-service/internal/services"
	"github.com/gin-gonic/gin"
)

// UserHandler contains handlers related to users
type UserHandlerImpl struct {
	userService services.UserService
}

// NewUserHandler creates a UserHandler instance
func NewUserHandler(userService services.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{userService: userService}
}

// GetUserByID handles a request to get a user by ID
func (u *UserHandlerImpl) GetUserByID(c *gin.Context) {}

// CreateUser handles a request to create a new user
func (u *UserHandlerImpl) CreateUser(c *gin.Context) {}
