package handlers

import (
	"github.com/gin-gonic/gin"
)

// UserHandler defines the interface for user-related handlers
type UserHandler interface {
	CreateUser(c *gin.Context)        // CreateUser handles the request to create a new user
	GetUserByID(c *gin.Context)       // GetUserByID handles the request to retrieve a user by ID
	UpdateUserProfile(c *gin.Context) // UpdateUserProfile handles the request to update a user
	ChangePassword(c *gin.Context)    // ChangePassword handles the request to change the user password
	DeleteUser(c *gin.Context)        // DeleteUser handles the request to delete a user
}
