package store

import "gihtub.com/tasks-hub/users-service/internal/entities"

// UserStore define storing operations for users
type UserStore interface {
	GetUserByID(userID int) (*entities.User, error)
	CreateUser(user *entities.User) error
}
