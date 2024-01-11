package handlers

import (
	service "gihtub.com/tasks-hub/users-service/internal/service"
	"github.com/gin-gonic/gin"
)

// UserHandler contains handlers related to users
type UserHandlerImpl struct {
	userService service.UserService
}

// NewUserHandler creates a UserHandler instance
func NewUserHandler(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{userService: userService}
}

// GetUserByID handles a request to retrieve a user by ID
func (u *UserHandlerImpl) GetUserByID(c *gin.Context) {}

// CreateUser handles a request to create a new user
func (u *UserHandlerImpl) CreateUser(c *gin.Context) {}
