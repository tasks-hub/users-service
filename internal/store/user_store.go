package store

import "gihtub.com/tasks-hub/users-service/internal/entities"

// UserStore define storing operations for users
type UserStore interface {
	GetUserByID(userID int) (*entities.User, error) // GetUserByID retrieves a user by ID from a database
	CreateUser(user *entities.User) error           // GetUserByID creates a user in a database
}
