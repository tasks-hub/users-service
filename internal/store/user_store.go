package store

import "github.com/tasks-hub/users-service/internal/entities"

// UserStore defines operations for user data storage
type UserStore interface {
	CreateUser(user *entities.CreateUserInput) (*entities.User, error)                // CreateUser creates a new user in the data store
	GetUserByID(userID string) (*entities.User, error)                                // GetUserByID retrieves a user from the data store based on the user ID
	GetUserByEmail(userCredentials *entities.UserCredentials) (*entities.User, error) // GetUserByEmail retrieves a user from the data store based on the user email and password
	UpdateUser(user *entities.User) error                                             // UpdateUser updates the information of an existing user in the data store
	DeleteUser(userID string) error                                                   // DeleteUser deletes a user from the data store based on the user ID
}
