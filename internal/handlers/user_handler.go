package handlers

import (
	"github.com/gin-gonic/gin"
)

// UserHandler defines the interface for user-related handlers
type UserHandler interface {
	GetUserByID(c *gin.Context) // GetUserByID handles the request to get a user by ID
	CreateUser(c *gin.Context)  // CreateUser handles the request to create a new user
}
