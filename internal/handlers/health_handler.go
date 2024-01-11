package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler defines the strict for health checks handler
type HealthHandler struct{}

// NewHealthHandler creates a HealthHandler instance
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health handles a request to check if the system is healthy
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
