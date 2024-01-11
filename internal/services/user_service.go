package services

import "gihtub.com/tasks-hub/users-service/internal/entities"

// UserService defines the interface for user-related services
type UserService interface {
	GetUserByID(userID int) (*entities.User, error) // GetUserByID retrieves a user by ID
	CreateUser(user *entities.User) error           // CreateUser creates a new user
}
