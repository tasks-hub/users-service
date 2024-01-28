package service

import "github.com/tasks-hub/users-service/internal/entities"

// UserService defines operations for the user service
type UserService interface {
	CreateUser(input entities.CreateUserInput) (*entities.User, error)
	GetUserByID(userID string) (*entities.User, error)
	GetUserByEmail(userCredentials *entities.UserCredentials) (*entities.User, error)
	UpdateUserProfile(userID string, input entities.UpdateUserInput) error
	ChangePassword(userID string, input entities.UpdateUserPasswordInput) error
	DeleteUser(userID string) error
}
