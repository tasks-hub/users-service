package store

import "github.com/tasks-hub/users-service/internal/entities"

// UserStore defines operations for user data storage
type UserStore interface {
	CreateUser(user *entities.CreateUserInput) error // CreateUser creates a new user in the data store
	GetUserByID(userID int) (*entities.User, error)  // GetUserByID retrieves a user from the data store based on the user ID
	UpdateUser(user *entities.User) error            // UpdateUser updates the information of an existing user in the data store
	DeleteUser(userID int) error                     // DeleteUser deletes a user from the data store based on the user ID
}
