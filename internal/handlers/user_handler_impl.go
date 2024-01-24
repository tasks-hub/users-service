// handlers/user_handler.go

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tasks-hub/users-service/internal/entities"
	service "github.com/tasks-hub/users-service/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandlerImpl contains handlers related to users
type UserHandlerImpl struct {
	userService service.UserService
}

// NewUserHandler creates a UserHandler instance
func NewUserHandler(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{userService: userService}
}

// CreateUser handles a request to create a new user
func (u *UserHandlerImpl) CreateUser(c *gin.Context) {
	var userInput entities.CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := u.userService.CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("user created successfully; ID: %s", userID)})
}

// GetUserByID handles a request to retrieve a user by ID
func (u *UserHandlerImpl) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := u.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile handles a request to update user profile
func (u *UserHandlerImpl) UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")

	var userProfileInput entities.UpdateUserInput
	if err := c.ShouldBindJSON(&userProfileInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.userService.UpdateUserProfile(userID, userProfileInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}

// ChangePassword handles a request to change user password
func (u *UserHandlerImpl) ChangePassword(c *gin.Context) {
	userID := c.Param("id")

	var passwordInput entities.UpdateUserInput
	if err := c.ShouldBindJSON(&passwordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.userService.ChangePassword(userID, passwordInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}

// DeleteUser handles a request to delete a user
func (u *UserHandlerImpl) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = u.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
