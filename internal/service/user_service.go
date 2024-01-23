package service

import "github.com/tasks-hub/users-service/internal/entities"

// UserService defines operations for the user service
type UserService interface {
	CreateUser(input entities.CreateUserInput) error
	GetUserByID(userID int) (*entities.User, error)
	UpdateUserProfile(userID int, input entities.UpdateUserInput) error
	ChangePassword(userID int, input entities.UpdateUserInput) error
	DeleteUser(userID int) error
}
