package services

import (
	"gihtub.com/tasks-hub/users-service/internal/entities"
	"gihtub.com/tasks-hub/users-service/internal/store"
)

// UserService contains business logic for users
type UserServiceImpl struct {
	userStore store.UserStore
}

// NewUserService creates a UserServiceImpl instance
func NewUserService(userStore store.UserStore) *UserServiceImpl {
	return &UserServiceImpl{userStore: userStore}
}

// GetUserByID gets a user by ID
func (u *UserServiceImpl) GetUserByID(userID int) (*entities.User, error) {
	return u.userStore.GetUserByID(userID)
}

// CreateUser creates a new user
func (u *UserServiceImpl) CreateUser(user *entities.User) error {
	return u.userStore.CreateUser(user)
}
