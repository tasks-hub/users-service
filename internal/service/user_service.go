package service

import "github.com/tasks-hub/users-service/internal/entities"

// UserService defines operations for the user service
type UserService interface {
	CreateUser(input entities.CreateUserInput) (string, error)
	GetUserByID(userID string) (*entities.User, error)
	UpdateUserProfile(userID string, input entities.UpdateUserInput) error
	ChangePassword(userID string, input entities.UpdateUserInput) error
	DeleteUser(userID int) error
}
